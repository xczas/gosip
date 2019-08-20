package manual

import (
	"fmt"

	"github.com/czas/gosip/auth/addin"
	u "github.com/czas/gosip/test/utils"
)

// ConfigReaderSpoAddinOnlyTest : test scenario
func ConfigReaderSpoAddinOnlyTest() {
	config := &addin.AuthCnfg{}
	err := config.ReadConfig(u.ResolveCnfgPath("./config/private.addin.json"))
	if err != nil {
		fmt.Printf("Error reading config: %v", err)
		return
	}
	fmt.Println(config)
}
