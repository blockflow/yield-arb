package protocols

import "fmt"

func GetProtocol(protocol string) (Protocol, error) {
	switch protocol {
	case "aavev2":
		return NewAaveV2Protocol(), nil
	// case "compound":
	// 	return NewCompoundProtocol(), nil
	// case "dydx":
	// 	return NewDydxProtocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
