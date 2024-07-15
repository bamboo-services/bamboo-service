/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBgYGC4aPY+ggEJiDFwMhSlFueXFiWn6pek5hbkJJak6ucmZuaEhrAyMK4+PiUtwJudA1kLzDAOhvUV+1EMk8dpmL5eemZJdmpqAdhUq+eRaYwMDP//w4xmxTD6Apo7VXEbDSLik/NTUvUySnIhzl55fErakbD//rfdRP7b1h81DV4SdkiF0Wv1mc4jGh5Tt5W9mrHsb0jsI++iVcnu3wSdlvv6/djk4JQoHSJ9muXqCRFFx2uh0jrxx+Sba7fVn9pvYMl3kG+GX9TSuEU39lW//2z3/u5x/eP3kiTkGe4zPKvvkTuwZUbslAIt3wNvj8gaiHH5G/XL1TOeeWu0VvZSaWuYdELX84PCp93Y/Cf0LJF55Na4L2zRfp2Lpq1h16s8oqIXntXNTb5sYWcmpZ720P/tfOcPJw9bXvfaHnvRrLtt05JtdVe/c796OXHfbh67Ogvbzy82SyboXd166dXuk2f4EqZkp4UanqsziwvdMje0dneeid+Wr+G3bnYF3rh2f0vGS+vuy2GtevlZ0fovTu/ferbtluXl7a+fewWuvva2evbMixYL8199OS5+a7vv24SbfZdXSp3y2r739ZrI7bGX+uNm335q0/Yo2K6zJWT73m96QskzZpw/fq/qQ0XVu9P+6Q6zUn7M6N+q19DupPWqft63aGGOIkN5BoYNki4KHksWqigLfa23Mz9957v1dJYZjoH1fItb12gI2nt0MDkydf4SN3C/FMj3tIDJcG+7YAOrj8ETSUvGShHe/ja9BRxqNl2sVdMZpLfIv1WdqRN3J1pZr6j7weqDi00nLXKYoXImoc6QaZkKz09BBWfHtBSVh0/a9W5b+Nh4ez7ac8CQc6L4Uo4fCpaJtXfSJDWWbhBgVPxeIlYkIxa07uf62ubbgdqVHawqvIZvlsozPgzQMJBRP/9FXN15Yc69ZtaGzy411Up7DVPUnkla8loWHmzibCzQmM0epMWX3+7m8N9aap/xG5VCjySVEw8rLS99LpTusFRRreQ27rXk47wefjjo4tSgmMO/7+3tT1t+7K5g3qnOBTaJCv3BTiwqExwmPNnonaz36LB0K4vRkbOHBXmOPgmcdIdn9l6/hUZ1fUoKNpY8tRNfVeRXaCqUJJvJJO3Z8n1bVmzLhgd7S88uVIzgOx0pp+fGIFaho6B2rrszXe73BeG7V2fFKbeu/Mj+PV39324RmVbZ2Tua3k/7ePdhqZxgk4rbuUMte1/HffXtUqpc78RXnVc7/5Zvk/TnHG0/j9fTeb/N1pXe8p8xYvXtYzU67cUKB50j+36xG+3OuNtn+rv1e7xZtpnsvpxXLwvuK594eCf0+b/dj7+cO5axcd1VK4UGI4HPjBdyw28kXEtduC3rVWakbe6lpNj7sdeuvQ7RnNEuujjwwnfuu2WtYddf+HWGfb9yWS3h8cbdOQZyzFGZV//Urb99r/Txzh8W8VrsutYbLp+R536iYJ+v+pJl95nN569a2Uw5v+V+OPsfn9TLa/4t3F4XFK1/NW3z7XUr1s7atilR6/eWLUvn5UWo997n+vXO6O6Ny3xx87VdF8csPmq1Vu7WcaGtV4Osrb+0824NfrtNO6xjc9++4MXu7GkfI3L/lq4TYd9i45XffSL2Z4PtT5PmFMllMiu/GrJN3+fStHfzkf3wYkk1JrzBmoWBIZoLUSwxMLCboxZLgtiKJXARxH8Cd8mJaQ4Hkjl4tDMyiTDjLschQIDhrSOIJliqwwxDlOOsSOU4DCxp9MVpGGapjmwqqAhHDkVVFFNX4DYVR4GO7H1s8QDz/n9He1YGvLGC2xwOFHPakcxB0s7KBgksVoY2RgaGY+BQAwQAAP//8QGi8GIHAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
