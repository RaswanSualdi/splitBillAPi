package controller

import (
	"SplitBillApi/payload/request"

	"github.com/gin-gonic/gin"
)

func (con *controller) SplitBill(c *gin.Context) {
	var input request.SplitBillRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(
			400,
			gin.H{
				"message": "payload error 1",
			},
		)

		return
	}

	splitBill, err := con.usecase.SplitBill(input)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "payload error",
		})
		return
	}

	c.JSON(201, splitBill)
}

func (con *controller) GetSplitBillByCustomerID(c *gin.Context) {

	param := c.Param("id")

	splitBills, err := con.usecase.GetSplitBillByCustomerID(param)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "payload error",
		})
		return
	}

	c.JSON(200, splitBills)

}
