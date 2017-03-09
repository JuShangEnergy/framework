package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/freeznet/tomato/config"
	"github.com/freeznet/tomato/rest"
)

// PublicController 处理密码修改与邮箱验证请求
type PublicController struct {
	beego.Controller
}

// VerifyEmail 处理验证邮箱请求
// 该接口从验证邮件内部发起请求，见 rest.SendVerificationEmail()
// @router /verify_email [get]
func (p *PublicController) VerifyEmail() {
	token := p.GetString("token")
	username := p.GetString("username")

	if config.TConfig.ServerURL == "" {
		p.missingPublicServerURL()
		return
	}

	if token == "" || username == "" {
		p.invalid()
		return
	}

	ok := rest.VerifyEmail(username, token)
	if ok {
		p.Ctx.Output.SetStatus(302)
		p.Ctx.Output.Header("location", config.TConfig.ServerURL+"/apps/verify_email_success?username="+username)
	} else {
		p.invalid()
	}
}

// ChangePassword 修改密码页面
// @router /choose_password [get]
func (p *PublicController) ChangePassword() {
	if config.TConfig.ServerURL == "" {
		p.missingPublicServerURL()
		return
	}

	data := strings.Replace(choosePasswordPage, "PARSE_SERVER_URL", `"`+config.TConfig.ServerURL+`"`, -1)
	p.Ctx.Output.Header("Content-Type", "text/html")
	p.Ctx.Output.Body([]byte(data))
}

// ResetPassword 处理实际的重置密码请求
// @router /request_password_reset [post]
func (p *PublicController) ResetPassword() {
	if config.TConfig.ServerURL == "" {
		p.missingPublicServerURL()
		return
	}

	username := p.GetString("username")
	token := p.GetString("token")
	newPassword := p.GetString("new_password")

	if token == "" || username == "" || newPassword == "" {
		p.invalid()
		return
	}

	err := rest.UpdatePassword(username, token, newPassword)
	if err == nil {
		p.Ctx.Output.SetStatus(302)
		p.Ctx.Output.Header("location", config.TConfig.ServerURL+"/apps/password_reset_success?username="+username)
	} else {
		p.Ctx.Output.SetStatus(302)
		location := config.TConfig.ServerURL + "/apps/choose_password"
		location += "?token=" + token
		location += "&id=" + config.TConfig.AppID
		location += "&username=" + username
		location += "&error=" + err.Error()
		location += "&app=" + config.TConfig.AppName
		p.Ctx.Output.Header("location", location)
	}
}

// RequestResetPassword 处理重置密码请求
// 该接口从重置密码邮件内部发起请求，见 rest.SendPasswordResetEmail()
// @router /request_password_reset [get]
func (p *PublicController) RequestResetPassword() {
	token := p.GetString("token")
	username := p.GetString("username")

	if config.TConfig.ServerURL == "" {
		p.missingPublicServerURL()
		return
	}

	if token == "" || username == "" {
		p.invalid()
		return
	}

	user := rest.CheckResetTokenValidity(username, token)
	if user != nil {
		p.Ctx.Output.SetStatus(302)
		location := config.TConfig.ServerURL + "/apps/choose_password"
		location += "?token=" + token
		location += "&id=" + config.TConfig.AppID
		location += "&username=" + username
		location += "&app=" + config.TConfig.AppName
		p.Ctx.Output.Header("location", location)
	} else {
		p.invalid()
	}
}

// InvalidLink 无效链接页面
// @router /invalid_link [get]
func (p *PublicController) InvalidLink() {
	p.Ctx.Output.Header("Content-Type", "text/html")
	p.Ctx.Output.Body([]byte(invalidLinkPage))
}

// PasswordResetSuccess 密码重置成功页面
// @router /password_reset_success [get]
func (p *PublicController) PasswordResetSuccess() {
	p.Ctx.Output.Header("Content-Type", "text/html")
	p.Ctx.Output.Body([]byte(passwordResetSuccessPage))
}

// VerifyEmailSuccess 验证邮箱成功页面
// @router /verify_email_success [get]
func (p *PublicController) VerifyEmailSuccess() {
	p.Ctx.Output.Header("Content-Type", "text/html")
	p.Ctx.Output.Body([]byte(verifyEmailSuccessPage))
}

func (p *PublicController) invalid() {
	p.Ctx.Output.SetStatus(302)
	p.Ctx.Output.Header("location", config.TConfig.ServerURL+"/apps/invalid_link")
}

func (p *PublicController) missingPublicServerURL() {
	p.Ctx.Output.SetStatus(404)
	p.Ctx.Output.Body([]byte("Not found."))
}

var verifyEmailSuccessPage = `
<!DOCTYPE html>
<html>
  <!-- This page is displayed whenever someone has successfully reset their password.
       Pro and Enterprise accounts may edit this page and tell Parse to use that custom
       version in their Parse app. See the App Settigns page for more information.
       This page will be called with the query param 'username'
   -->
  <head>
  <title>Email Verification</title>
  <style type='text/css'>
    h1 {
      color: #0067AB;
      display: block;
      font: inherit;
      font-family: 'Open Sans', 'Helvetica Neue', Helvetica;
      font-size: 30px;
      font-weight: 600;
      height: 30px;
      line-height: 30px;
      margin: 45px 0px 0px 45px;
      padding: 0px 8px 0px 8px;
    }
  </style>
  <body>
    <h1>Successfully verified your email!</h1>
  </body>
</html>
`

var passwordResetSuccessPage = `
<!DOCTYPE html>
<html>
  <!-- This page is displayed whenever someone has successfully reset their password.
       Pro and Enterprise accounts may edit this page and tell Parse to use that custom
       version in their Parse app. See the App Settigns page for more information.
       This page will be called with the query param 'username'
    -->
  <head>
  <title>Password Reset</title>
  <style type='text/css'>
    h1 {
      color: #0067AB;
      display: block;
      font: inherit;
      font-family: 'Open Sans', 'Helvetica Neue', Helvetica;
      font-size: 30px;
      font-weight: 600;
      height: 30px;
      line-height: 30px;
      margin: 45px 0px 0px 45px;
      padding: 0px 8px 0px 8px;
    }
  </style>
  <body>
    <h1>Successfully updated your password!</h1>
  </body>
</html>
`

var invalidLinkPage = `
<!DOCTYPE html>
<!-- This page is displayed when someone navigates to a verify email or reset password link
     but their security token is wrong. This can either mean the user has clicked on a
     stale link (i.e. re-click on a password reset link after resetting their password) or
     (rarely) this could be a sign of a malicious user trying to tamper with your app.
 -->
<html>
  <head>
  <title>Invalid Link</title>
  <style type='text/css'>
    .container {
      border-width: 0px;
      display: block;
      font: inherit;
      font-family: 'Helvetica Neue', Helvetica;
      font-size: 16px;
      height: 30px;
      line-height: 16px;
      margin: 45px 0px 0px 45px;
      padding: 0px 8px 0px 8px;
      position: relative;
      vertical-align: baseline;
    }

    h1, h2, h3, h4, h5 {
      color: #0067AB;
      display: block;
      font: inherit;
      font-family: 'Open Sans', 'Helvetica Neue', Helvetica;
      font-size: 30px;
      font-weight: 600;
      height: 30px;
      line-height: 30px;
      margin: 0 0 15px 0;
      padding: 0 0 0 0;
    }
  </style>
  <body> 
    <div class="container">
      <h1>Invalid Link</h1>
    </div> 
  </body>
</html>
`

var choosePasswordPage = `
<!DOCTYPE html>
<html>
  <!-- This page is displayed when someone clicks a valid 'reset password' link.
       Users should feel free to add to this page (i.e. branding or security widgets)
       but should be sure not to delete any of the form inputs or the javascript from the
       template file. This javascript is what adds the necessary values to authenticate
       this session with Parse.
       The query params 'username' and 'app' hold the friendly names for your current user and
       your app. You should feel free to incorporate their values to make the page more personal.
       If you are missing form parameters in your POST, Parse will navigate back to this page and
       add an 'error' query parameter.
  -->
  <head>
  <title>Password Reset</title>
  <style type='text/css'>
    h1 {
      display: block;
      font: inherit;
      font-size: 30px;
      font-weight: 600;
      height: 30px;
      line-height: 30px;
      margin: 45px 0px 45px 0px;
      padding: 0px 8px 0px 8px;
    }

    .error {
      color: red;
      padding: 0px 8px 0px 8px;
      margin: -25px 0px -20px 0px;
    }

    body {
      font-family: 'Open Sans', 'Helvetica Neue', Helvetica;
      color: #0067AB;
      margin: 15px 99px 0px 98px;
    }

    label {
      color: #666666;
    }
    form {
      margin: 0px 0px 45px 0px;
      padding: 0px 8px 0px 8px;
    }
    form > * {
      display: block;
      margin-top: 25px;
      margin-bottom: 7px;
    }

    button {
      font-size: 22px;
      color: white;
      background: #0067AB;
      -moz-border-radius: 5px;
      -webkit-border-radius: 5px;
      -o-border-radius: 5px;
      -ms-border-radius: 5px;
      -khtml-border-radius: 5px;
      border-radius: 5px;
      background-image: -webkit-gradient(linear,50% 0,50% 100%,color-stop(0%,#0070BA),color-stop(100%,#00558C));
      background-image: -webkit-linear-gradient(#0070BA,#00558C);
      background-image: -moz-linear-gradient(#0070BA,#00558C);
      background-image: -o-linear-gradient(#0070BA,#00558C);
      background-image: -ms-linear-gradient(#0070BA,#00558C);
      background-image: linear-gradient(#0070BA,#00558C);
      -moz-box-shadow: inset 0 1px 0 0 #0076c4;
      -webkit-box-shadow: inset 0 1px 0 0 #0076c4;
      -o-box-shadow: inset 0 1px 0 0 #0076c4;
      box-shadow: inset 0 1px 0 0 #0076c4;
      border: 1px solid #005E9C;
      padding: 10px 14px;
      cursor: pointer;
      outline: none;
      display: block;
      font-family: "Helvetica Neue",Helvetica;

      -webkit-box-align: center;
      text-align: center;
      box-sizing: border-box;
      letter-spacing: normal;
      word-spacing: normal;
      line-height: normal;
      text-transform: none;
      text-indent: 0px;
      text-shadow: none;
    }

    button:hover {
      background-image: -webkit-gradient(linear,50% 0,50% 100%,color-stop(0%,#0079CA),color-stop(100%,#005E9C));
      background-image: -webkit-linear-gradient(#0079CA,#005E9C);
      background-image: -moz-linear-gradient(#0079CA,#005E9C);
      background-image: -o-linear-gradient(#0079CA,#005E9C);
      background-image: -ms-linear-gradient(#0079CA,#005E9C);
      background-image: linear-gradient(#0079CA,#005E9C);
      -moz-box-shadow: inset 0 0 0 0 #0076c4;
      -webkit-box-shadow: inset 0 0 0 0 #0076c4;
      -o-box-shadow: inset 0 0 0 0 #0076c4;
      box-shadow: inset 0 0 0 0 #0076c4;
    }

    button:active {
      background-image: -webkit-gradient(linear,50% 0,50% 100%,color-stop(0%,#00395E),color-stop(100%,#005891));
      background-image: -webkit-linear-gradient(#00395E,#005891);
      background-image: -moz-linear-gradient(#00395E,#005891);
      background-image: -o-linear-gradient(#00395E,#005891);
      background-image: -ms-linear-gradient(#00395E,#005891);
      background-image: linear-gradient(#00395E,#005891);
    }

    input {
      color: black;
      cursor: auto;
      display: inline-block;
      font-family: 'Helvetica Neue', Helvetica;
      font-size: 25px;
      height: 30px;
      letter-spacing: normal;
      line-height: normal;
      margin: 2px 0px 2px 0px;
      padding: 5px;
      text-transform: none;
      vertical-align: baseline;
      width: 500px;
      word-spacing: 0px; 
    }

  </style>
</head>
<body>
  <h1>Reset Your Password<span id='app'></span></h1>
  <noscript>We apologize, but resetting your password requires javascript</noscript>
  <div class='error' id='error'></div>
  <form id='form' action='#' method='POST'>
    <label>New Password for <span id='username_label'></span></label>
    <input name="new_password" type="password" />
    <input name='utf-8' type='hidden' value='✓' />
    <input name="username" id="username" type="hidden" />
    <input name="token" id="token" type="hidden" />
    <button>Change Password</button>
  </form>

<script language='javascript' type='text/javascript'>
  window.onload = function() {
    var urlParams = {};
    (function () {
        var pair, // Really a match. Index 0 is the full match; 1 & 2 are the key & val.
            tokenize = /([^&=]+)=?([^&]*)/g,
            // decodeURIComponents escapes everything but will leave +s that should be ' '
            re_space = function (s) { return decodeURIComponent(s.replace(/\+/g, " ")); },
            // Substring to cut off the leading '?'
            querystring = window.location.search.substring(1);

        while (pair = tokenize.exec(querystring))
           urlParams[re_space(pair[1])] = re_space(pair[2]);
    })();

    var id = urlParams['id'];
    var base = PARSE_SERVER_URL;
    document.getElementById('form').setAttribute('action', base + '/apps' + '/request_password_reset');
    document.getElementById('username').value = urlParams['username'];
    document.getElementById('username_label').appendChild(document.createTextNode(urlParams['username']));

    document.getElementById('token').value = urlParams['token'];
    if (urlParams['error']) {
      document.getElementById('error').appendChild(document.createTextNode(urlParams['error']));
    }
    if (urlParams['app']) {
      document.getElementById('app').appendChild(document.createTextNode(' for ' + urlParams['app']));
    }
  }
</script>
</body>
`
