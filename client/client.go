package client

import (
	"fmt"	
)

func wrapeError(customError string, orginalError error) error {
	return fmt.Errorf("%s %v", customError, orginalError)
}
