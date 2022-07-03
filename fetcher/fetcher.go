package fetcher

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
)

//根据指定url获取响应数据
func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	//增加header选项
	reqest.Header.Add("Cookie", "_uab_collina=165624418601871554304163; adv=ad_logid_url%3Dhttps%253A%252F%252Ftrace.51job.com%252Ftrace.php%253Fpartner%253Dsem_pcbaidu5_153404%2526ajp%253DaHR0cHM6Ly9ta3QuNTFqb2IuY29tL3RnL3NlbS9MUF8yMDIwXzEuaHRtbD9mcm9tPWJhaWR1YWQ%253D%2526k%253Dd946ba049bfb67b64f408966cbda3ee9%2526bd_vid%253D12354641525106228121%26%7C%26; guid=9200419b82ab93f16ccb7b6f84c77cd0; _ujz=MTQ4NTAwMDU2MA%3D%3D; ps=needv%3D0; slife=lowbrowser%3Dnot%26%7C%26lastlogindate%3D20220626%26%7C%26securetime%3DBTkEMVEyWTsCYgE7AT5aMFdmAj0%253D; nsearch=jobarea%3D%26%7C%26ord_field%3D%26%7C%26recentSearch0%3D%26%7C%26recentSearch1%3D%26%7C%26recentSearch2%3D%26%7C%26recentSearch3%3D%26%7C%26recentSearch4%3D%26%7C%26collapse_expansion%3D; sajssdk_2015_cross_new_user=1; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%229200419b82ab93f16ccb7b6f84c77cd0%22%2C%22first_id%22%3A%221819fd9ae84e5f-08d4988be78f108-1c525635-1296000-1819fd9ae85255%22%2C%22props%22%3A%7B%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTgxOWZkOWFlODRlNWYtMDhkNDk4OGJlNzhmMTA4LTFjNTI1NjM1LTEyOTYwMDAtMTgxOWZkOWFlODUyNTUiLCIkaWRlbnRpdHlfbG9naW5faWQiOiI5MjAwNDE5YjgyYWI5M2YxNmNjYjdiNmY4NGM3N2NkMCJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%229200419b82ab93f16ccb7b6f84c77cd0%22%7D%2C%22%24device_id%22%3A%221819fd9ae84e5f-08d4988be78f108-1c525635-1296000-1819fd9ae85255%22%7D; partner=localhost; acw_tc=2f624a3b16562462138772468e415aa59d394b3970a3ff86427a675dfad76a; acw_sc__v2=62b850135053bae63978619d04687db098303e08; search=jobarea%7E%60000000%7C%21ord_field%7E%600%7C%21recentSearch0%7E%60000000%A1%FB%A1%FA000000%A1%FB%A1%FA0000%A1%FB%A1%FA00%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA9%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA0%A1%FB%A1%FAJava%A1%FB%A1%FA2%A1%FB%A1%FA1%7C%21recentSearch1%7E%60000000%A1%FB%A1%FA000000%A1%FB%A1%FA0000%A1%FB%A1%FA00%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA9%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA0%A1%FB%A1%FAJava%BF%AA%B7%A2%B9%A4%B3%CC%CA%A6%A1%FB%A1%FA2%A1%FB%A1%FA1%7C%21recentSearch2%7E%60000000%A1%FB%A1%FA000000%A1%FB%A1%FA0000%A1%FB%A1%FA00%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA9%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA0%A1%FB%A1%FAgo%A1%FB%A1%FA2%A1%FB%A1%FA1%7C%21recentSearch3%7E%60000000%A1%FB%A1%FA000000%A1%FB%A1%FA0000%A1%FB%A1%FA00%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA99%A1%FB%A1%FA9%A1%FB%A1%FA99%A1%FB%A1%FA%A1%FB%A1%FA0%A1%FB%A1%FAjava%A1%FB%A1%FA2%A1%FB%A1%FA1%7C%21; ssxmod_itna=QqUxcD2QD=KrqGHdWexj2CxRjUDC9UlQj/jDBqEeiNDnD8x7YDv+GH/DGo8TsuGUNeiPW4+Pdnh4itSF/O2KWG/G8uDWli0qiTD4q07Db4GkDAqiOD7dRoufHqDPDbxYPDGQKGCDbq0CRxDa1p5Di6KDcZ5zDCnzsxDCUnH7QIhxAD5kSaF7DBbeqjnAr0uKvmE4YlxT3=0YZlxjr1xxsC4Ttlxbkj+PI4jtKfiN+3D=; ssxmod_itna2=QqUxcD2QD=KrqGHdWexj2CxRjUDC9UlQj/D8MCqgqGXieGa9Q9IOka+x8x2KY5D0jAa07K5nfyDzd+OigKIMUewiWkMYeMvbekyO0fRfMDS=L6LMC9dId6l74HjMUjgE/eG2ijBstKr03uenF/xY4RQAQk77QYenNDjKD+PWrsODWxE7i4D=")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	if err != nil {
		return nil, err
	}
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()
	//状态码判断
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error:status code :%v", resp.StatusCode)
		return nil, fmt.Errorf("get incorrect")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//解决中文乱码
	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	return []byte(bodystr), nil
}
