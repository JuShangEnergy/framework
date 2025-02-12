package files

import (
	"github.com/JuShangEnergy/framework/cloud"
	"github.com/JuShangEnergy/framework/config"
	"github.com/JuShangEnergy/framework/types"
	"github.com/JuShangEnergy/framework/utils"
)

var adapter filesAdapter

// init 初始化文件处理模块
// 当前支持本地文件存储模块、数据库文件存储
// 后续可增加第三方网络文件存储模块
func init() {
	a := config.TConfig.FileAdapter
	if a == "Disk" {
		adapter = newFileSystemAdapter(config.TConfig.AppID)
	} else if a == "GridFS" {
		adapter = newGridStoreAdapter()
	} else if a == "Qiniu" {
		adapter = newQiniuAdapter()
	} else if a == "Sina" {
		adapter = newSinaAdapter()
	} else if a == "Tencent" {
		adapter = newTencentAdapter()
	} else if a == "HDFS" {
		adapter = newHDFSAdapter(config.TConfig.HDFSNameNode, config.TConfig.HDFSUser, config.TConfig.HDFSRoot)
	} else {
		adapter = newFileSystemAdapter(config.TConfig.AppID)
	}
}

// GetFileData 获取文件数据
func GetFileData(filename string) ([]byte, error) {
	return adapter.getFileData(filename)
}

// CreateFile 创建文件，返回文件地址与文件名
func CreateFile(filename string, data []byte, contentType string, user types.M, info *types.RequestInfo) types.M {
	extname := utils.ExtName(filename)
	if extname == "" && contentType != "" && utils.LookupExtension(contentType) != "" {
		filename = filename + "." + utils.LookupExtension(contentType)
	} else if extname != "" && contentType == "" {
		contentType = utils.LookupContentType(filename)
	}
	orifilename := filename
	filename = utils.CreateFileName() + "-" + filename
	location := adapter.getFileLocation(filename)

	//hasBeforeFileUploadHook := cloud.TriggerExists(cloud.TypeBeforeFileUpload, extname)
	//if hasBeforeFileUploadHook == false {
	//	return nil
	//}
	response, err := maybeRunTrigger(cloud.TypeBeforeFileUpload, extname, orifilename, data, contentType, user, info)
	if err != nil {
		return nil
	}

	if response != nil && response["data"] != nil {
		newdata, ok := response["data"].([]byte)
		if ok {
			err = adapter.createFile(filename, newdata, contentType)
		} else {
			err = adapter.createFile(filename, data, contentType)
		}
	} else {
		err = adapter.createFile(filename, data, contentType)
	}

	if err != nil {
		return nil
	}

	//hasAfterFileUploadHook := cloud.TriggerExists(cloud.TypeAfterFileUpload, extname)
	//if hasAfterFileUploadHook == false {
	//	return nil
	//}
	result, err := maybeRunAfterTrigger(cloud.TypeAfterFileUpload, extname, orifilename, filename, location, info)
	if err != nil {
		return nil
	}

	return types.M{
		"url":  location,
		"name": filename,
		"data": result["data"],
	}

}

// DeleteFile 删除文件
func DeleteFile(filename string) error {
	return adapter.deleteFile(filename)
}

// ExpandFilesInObject 展开文件对象
// 展开之后的文件对象如下
//
//	{
//		"__type": "File",
//		"url": "http://example.com/pic.jpg",
//		"name": "pic.jpg",
//	}
func ExpandFilesInObject(object interface{}) {
	if object == nil {
		return
	}
	if objs := utils.A(object); objs != nil {
		for _, obj := range objs {
			ExpandFilesInObject(obj)
		}
	}

	obj := utils.M(object)
	if obj == nil {
		return
	}

	for _, v := range obj {
		fileObject := utils.M(v)
		if fileObject != nil && fileObject["__type"] == "File" {
			if fileObject["url"] != nil {
				continue
			}
			filename := utils.S(fileObject["name"])
			fileObject["url"] = adapter.getFileLocation(filename)
		}
	}
}

// GetFileStream 获取文件流
func GetFileStream(filename string) (FileStream, error) {
	return adapter.getFileStream(filename)
}

// GetAdapterName ...
func GetAdapterName() string {
	return adapter.getAdapterName()
}

// filesAdapter 规定了文件存储模块需要实现的接口
type filesAdapter interface {
	createFile(filename string, data []byte, contentType string) error
	deleteFile(filename string) error
	getFileData(filename string) ([]byte, error)
	getFileLocation(filename string) string
	getFileStream(filename string) (FileStream, error)
	getAdapterName() string
}

// FileStream 规定了文件流需要实现的接口
type FileStream interface {
	Seek(offset int64, whence int) (ret int64, err error)
	Read(b []byte) (n int, err error)
	Size() (bytes int64)
	Close() (err error)
}
