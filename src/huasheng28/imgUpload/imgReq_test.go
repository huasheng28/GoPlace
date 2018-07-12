package imgUpload

import "testing"

func TestDoUpload(t *testing.T) {
	areq:=ARequest{
		Headers:map[string]string{"X-Api-Key":"JDZBNH3KZTYBT8ZZUO84R7INN7HRHIRJ",},
		Url:"https://api-intl.seeunsee.cn/v1/identify",
		ImgParam:"image",
		OutTime:30,
		ImgFolder:"",
		Params:map[string]string{
				"app_name":"maxigenes-app",
				"ip":"14.123.254.166",
				"dev_os":"Android",
				"dev_os_ver":"6.0",
				"dev_model":"OPPO R9sk",
				"latitude":"23.137466",
				"longitude":"113.352425",
				"userid":"10021",
				"nickname":"测试用户",
				"sex":"1",
				"headimgurl":"http://wx.qlogo.cn/mmopen/vi_32/XwVzYTewdIXexzpXW5QfpOBKOMRdFm3BxdLS7YKWxGKGCocwtib25QJAhNMpxiaGvsWO9um3Uexyw4q0wKCK8iciag/0",
			},
	}
	imgPath:="F:/Code/GoPlace/cleanupload/image/1.jpg"

	t.Log("前置条件完成")
	{
		t.Logf("测试 %v 接口",areq.Url)
		{
			respBody:=DoUpload(areq,imgPath)
			if respBody != ""{
				t.Logf("返回： %v",respBody)
			}else {
				t.Log("返回错误")
			}
		}
	}

}