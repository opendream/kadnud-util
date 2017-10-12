package util

import (
	"os"
	"fmt"
	"log"
	"strings"
	"github.com/opendream/deeperror"
	"strconv"
	"github.com/gin-gonic/gin"
	jwtgo "gopkg.in/dgrijalva/jwt-go.v3"
)

var Provinces = map[string]string{
	"TH-37": "อำนาจเจริญ",
	"TH-15": "อ่างทอง",
	"TH-14": "พระนครศรีอยุธยา",
	"TH-10": "กรุงเทพมหานคร",
	"TH-38": "บึงกาฬ",
	"TH-31": "บุรีรัมย์",
	"TH-24": "ฉะเชิงเทรา",
	"TH-18": "ชัยนาท",
	"TH-36": "ชัยภูมิ",
	"TH-22": "จันทบุรี",
	"TH-50": "เชียงใหม่",
	"TH-57": "เชียงราย",
	"TH-20": "ชลบุรี",
	"TH-86": "ชุมพร",
	"TH-46": "กาฬสินธุ์",
	"TH-62": "กำแพงเพชร",
	"TH-71": "กาญจนบุรี",
	"TH-40": "ขอนแก่น",
	"TH-81": "กระบี่",
	"TH-52": "ลำปาง",
	"TH-51": "ลำพูน",
	"TH-42": "เลย",
	"TH-16": "ลพบุรี",
	"TH-58": "แม่ฮ่องสอน",
	"TH-44": "มหาสารคาม",
	"TH-49": "มุกดาหาร",
	"TH-26": "นครนายก",
	"TH-73": "นครปฐม",
	"TH-48": "นครพนม",
	"TH-30": "นครราชสีมา",
	"TH-60": "นครสวรรค์",
	"TH-80": "นครศรีธรรมราช",
	"TH-55": "น่าน",
	"TH-96": "นราธิวาส",
	"TH-39": "หนองบัวลำภู",
	"TH-43": "หนองคาย",
	"TH-12": "นนทบุรี",
	"TH-13": "ปทุมธานี",
	"TH-94": "ปัตตานี",
	"TH-82": "พังงา",
	"TH-93": "พัทลุง",
	"TH-56": "พะเยา",
	"TH-67": "เพชรบูรณ์",
	"TH-76": "เพชรบุรี",
	"TH-66": "พิจิตร",
	"TH-65": "พิษณุโลก",
	"TH-54": "แพร่",
	"TH-83": "ภูเก็ต",
	"TH-25": "ปราจีนบุรี",
	"TH-77": "ประจวบคีรีขันธ์",
	"TH-85": "ระนอง",
	"TH-70": "ราชบุรี",
	"TH-21": "ระยอง",
	"TH-45": "ร้อยเอ็ด",
	"TH-27": "สระแก้ว",
	"TH-47": "สกลนคร",
	"TH-11": "สมุทรปราการ",
	"TH-74": "สมุทรสาคร",
	"TH-75": "สมุทรสงคราม",
	"TH-19": "สระบุรี",
	"TH-91": "สตูล",
	"TH-17": "สิงห์บุรี",
	"TH-33": "ศรีสะเกษ",
	"TH-90": "สงขลา",
	"TH-64": "สุโขทัย",
	"TH-72": "สุพรรณบุรี",
	"TH-84": "สุราษฎร์ธานี",
	"TH-32": "สุรินทร์",
	"TH-63": "ตาก",
	"TH-92": "ตรัง",
	"TH-23": "ตราด",
	"TH-34": "อุบลราชธานี",
	"TH-41": "อุดรธานี",
	"TH-61": "อุทัยธานี",
	"TH-53": "อุตรดิตถ์",
	"TH-95": "ยะลา",
	"TH-35": "ยโสธร",
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetenvByDomain(key, domain, fallback string) string {
	value := os.Getenv(fmt.Sprintf("%v__%v", key, domain))
	if len(value) == 0 {
		return fallback
	}
	return value
}

func Elog(num int64, msg string, err error, domain string, oid string, mid string, email string) {
	debug := Getenv("DEBUG", "false") == "true"

	msg = deeperror.New(num, msg, err).Error()
	if !debug {
		msg = strings.Replace(msg, "\n", "\r", -1)
		log.Println(fmt.Sprintf("%v \r [domain: %v][oid: %v][mid: %v][email: %v]",
			msg, domain, oid, mid, email,
		))
	} else {
		log.Println(fmt.Sprintf("%v \n [domain: %v][oid: %v][mid: %v][email: %v]",
			msg, domain, oid, mid, email,
		))
	}
}

func Plog(msg string, domain string, oid string, mid string, email string) {
	debug := Getenv("DEBUG", "false") == "true"

	if !debug {
		msg = strings.Replace(msg, "\n", "\r", -1)
		fmt.Println(fmt.Sprintf("Message: %v \r [domain: %v][oid: %v][mid: %v][email: %v]",
			msg, domain, oid, mid, email,
		))
	} else {
		fmt.Println(fmt.Sprintf("Message: %v \n [domain: %v][oid: %v][mid: %v][email: %v]",
			msg, domain, oid, mid, email,
		))
	}
}

func TzToSeconds(tz string) (seconds int) {
	if len(tz) != 5 {
		return
	}
	number, err := strconv.Atoi(tz[1:5])
	if err != nil {
		return
	}

	seconds = int((float64(number)/100) * 3600)
	if tz[0:1] == "-" {
		seconds = -seconds
	}
	return
}

func GetJwtPayload(c *gin.Context) (jwtgo.MapClaims, error) {
	if jwtPayload, found := c.Get("JWT_PAYLOAD"); found {
		claims := jwtPayload.(jwtgo.MapClaims)
		return claims, nil
	}

	err := fmt.Errorf("Cannot load JWT")
	return nil, err
}

func GetMemberIdFromJwtPayload(c *gin.Context) (string, error) {
	// assign member_id
	if jwtPayload, found := c.Get("JWT_PAYLOAD"); found {
		claims := jwtPayload.(jwtgo.MapClaims)

		return claims["id"].(string), nil
	}

	err := fmt.Errorf("Cannot load JWT")
	return "", err
}

func GetDomainFromJwtPayload(c *gin.Context) (string, error) {
	// assign member_id
	if jwtPayload, found := c.Get("JWT_PAYLOAD"); found {
		claims := jwtPayload.(jwtgo.MapClaims)

		return claims["domain"].(string), nil
	}

	err := fmt.Errorf("Cannot load JWT")
	return "", err
}

func GetProvinceName(code string) string {
	return Provinces[code]
}
