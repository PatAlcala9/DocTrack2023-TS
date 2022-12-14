package main

import (
	"database/sql"
	"net/http"
  // "html/template"
  // "encoding/json"

  "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// DEV
var connection string = "root:superuser@tcp(localhost:3306)/ocbodoctracksys"

// SERVER
// var connection string = "iips:iipsuser@tcp(192.168.7.100:3306)/iips"

func main() {
	connect()
}

func connect() {
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := gin.Default()
  router.Use(cors.Default())
  router.LoadHTMLGlob("templates/*")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
      "content": "This is an index page...",
    })
  })

  router.GET("/api/:method", func(c *gin.Context) {
    var result string
    method := c.Param("method")

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    if method == "CheckConnection" {
      err = db.QueryRow("SELECT 1 AS result").Scan(&result)
      if err != nil {
        panic(err.Error())
      }
      c.JSON(http.StatusOK, gin.H{
        "result": result,
      })

    } else if method == "GetIncoming" {
      var result2, result3, result4 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}

      results, err := db.Query("SELECT IFNULL(entryCodeNo, '') AS result, IFNULL(DATE_FORMAT(receivedDate, '%M %e, %Y'), '') AS result2, IFNULL(sourceName, '') AS result3, IFNULL(subjectInfo, '') AS result4 FROM incoming LIMIT 100")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
			})

    } else if method == "GetOutgoing" {
      var result2, result3, result4, result5 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}

      results, err := db.Query("SELECT IFNULL(referenceNo, '') AS result, IFNULL(respondent, '') AS result2, IFNULL(address, '') AS result3, IFNULL(subjectInfo, '') AS result4, (DATE_FORMAT(date_released, '%M %e, %Y')) AS result5 FROM outgoing WHERE date_released IS NOT NULL AND date_released <> 'NULL' LIMIT 100")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4, &result5)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
        array5 = append(array5, result5)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
        "result5": array5,
			})
    }
  })




  router.GET("/api/:method/:data", func(c *gin.Context) {
    var result string
    method := c.Param("method")
    data := c.Param("data")

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    if method == "GetIncomingByEntryCode" {
      var result2, result3, result4 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}

      results, err := db.Query("SELECT IFNULL(entryCodeNo, '') AS result, IFNULL(DATE_FORMAT(receivedDate, '%M %e, %Y'), '') AS result2, IFNULL(sourceName, '') AS result3, IFNULL(subjectInfo, '') AS result4 FROM incoming WHERE entryCodeNo = ?", data)
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
			})

    } else if method == "GetIncomingBySourceName" {
      var result2, result3, result4 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}

      results, err := db.Query("SELECT IFNULL(entryCodeNo, '') AS result, IFNULL(DATE_FORMAT(receivedDate, '%M %e, %Y'), '') AS result2, IFNULL(sourceName, '') AS result3, IFNULL(subjectInfo, '') AS result4 FROM incoming WHERE sourceName LIKE ?", "%" + data + "%")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
			})

    } else if method == "GetIncomingBySubject" {
      var result2, result3, result4 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}

      results, err := db.Query("SELECT IFNULL(entryCodeNo, '') AS result, IFNULL(DATE_FORMAT(receivedDate, '%M %e, %Y'), '') AS result2, IFNULL(sourceName, '') AS result3, IFNULL(subjectInfo, '') AS result4 FROM incoming WHERE subjectInfo LIKE ?", "%" + data + "%")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
			})

    } else if method == "GetIncomingDetails" {
      var result2, result3 string

      err = db.QueryRow("SELECT IFNULL(subDetails, '') AS result, IFNULL(attachments, '') AS result2, IFNULL(bo_notes, '') AS result3 FROM incoming WHERE entryCodeNo = ?", data).Scan(&result, &result2, &result3)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
        "result2": result2,
        "result3": result3,
			})

    } else if method == "GetIncomingDocLog" {
      var result2 string
      array := []string{}
      array2 := []string{}

      results, err := db.Query("SELECT IFNULL(DATE_FORMAT(date_forwarded, '%M %e, %Y'), '') AS result, IFNULL(forwardedto_name, '') AS result2 FROM doc_logs WHERE entryCodeNo = ?", data)
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
			})

    } else if method == "GetIncomingActionLog" {
      var result2 string
      array := []string{}
      array2 := []string{}

      results, err := db.Query("SELECT DISTINCT IFNULL(DATE_FORMAT(action_date, '%M %e, %Y'), '') AS result, IFNULL(actionMade, '') AS result2 FROM action_logs WHERE entryCodeNo = ?", data)
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
			})

    } else if method == "GetOutgoingByReference" {
      var result2, result3, result4, result5 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}

      results, err := db.Query("SELECT IFNULL(referenceNo, '') AS result, IFNULL(respondent, '') AS result2, IFNULL(address, '') AS result3, IFNULL(subjectInfo, '') AS result4, (DATE_FORMAT(date_released, '%M %e, %Y')) AS result5 FROM outgoing WHERE date_released IS NOT NULL AND date_released <> 'NULL' AND referenceNo LIKE ?", "8751-04-%" + data)
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4, &result5)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
        array5 = append(array5, result5)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
        "result5": array5,
			})

    } else if method == "GetOutgoingByRespondent" {
      var result2, result3, result4, result5 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}

      results, err := db.Query("SELECT IFNULL(referenceNo, '') AS result, IFNULL(respondent, '') AS result2, IFNULL(address, '') AS result3, IFNULL(subjectInfo, '') AS result4, (DATE_FORMAT(date_released, '%M %e, %Y')) AS result5 FROM outgoing WHERE date_released IS NOT NULL AND date_released <> 'NULL' AND respondent LIKE ?", "%" + data + "%")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4, &result5)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
        array5 = append(array5, result5)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
        "result5": array5,
			})

    } else if method == "GetOutgoingByAddress" {
      var result2, result3, result4, result5 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}

      results, err := db.Query("SELECT IFNULL(referenceNo, '') AS result, IFNULL(respondent, '') AS result2, IFNULL(address, '') AS result3, IFNULL(subjectInfo, '') AS result4, (DATE_FORMAT(date_released, '%M %e, %Y')) AS result5 FROM outgoing WHERE date_released IS NOT NULL AND date_released <> 'NULL' AND address LIKE ?", "%" + data + "%")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4, &result5)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
        array5 = append(array5, result5)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
        "result5": array5,
			})

    } else if method == "GetOutgoingBySubject" {
      var result2, result3, result4, result5 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}

      results, err := db.Query("SELECT IFNULL(referenceNo, '') AS result, IFNULL(respondent, '') AS result2, IFNULL(address, '') AS result3, IFNULL(subjectInfo, '') AS result4, (DATE_FORMAT(date_released, '%M %e, %Y')) AS result5 FROM outgoing WHERE date_released IS NOT NULL AND date_released <> 'NULL' AND subjectInfo LIKE ?", "%" + data + "%")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4, &result5)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
        array5 = append(array5, result5)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
        "result5": array5,
			})

    } else if method == "GetOutgoingDetails" {
      var result2, result3 string

      err = db.QueryRow("SELECT IFNULL(sender_name, '') AS result, IFNULL(subjectDetails, '') AS result2, IFNULL(Attachments, '') AS result3 FROM outgoing WHERE referenceNo = ?", "8751-04-" + data).Scan(&result, &result2, &result3)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
        "result2": result2,
        "result3": result3,
			})

    } else if method == "GetOutgoingDocLog" {
      var result2 string
      array := []string{}
      array2 := []string{}

      results, err := db.Query("SELECT IFNULL(DATE_FORMAT(date_forwarded, '%M %e, %Y'), '') AS result, IFNULL(UPPER(forwardedto_name), '') AS result2 FROM doc_logs WHERE referenceNo = ? ORDER BY date_forwarded", "8751-04-" + data)
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
			})

    } else if method == "GetOutgoingActionLog" {
      var result2 string
      array := []string{}
      array2 := []string{}

      results, err := db.Query("SELECT IFNULL(DATE_FORMAT(action_date, '%M %e, %Y'), '') AS result, IFNULL(actionMade, '') AS result2 FROM action_logs WHERE referenceNo = ? ORDER BY action_date", "8751-04-" + data)
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
			})

    } else if method == "CheckUsername" {
      err = db.QueryRow("SELECT COUNT(userid) AS result FROM user WHERE username = ?", data).Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})

    } else if method == "CheckPassword" {
      err = db.QueryRow("SELECT password AS result FROM user WHERE username = ?", data).Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})

    } else if method == "GetUserDetails" {
      var result2, result3, result4, result5, result6, result7 string

      err = db.QueryRow("SELECT userid AS result, employeeName AS result2, is_incoming AS result3, is_outgoing AS result4, is_releasing AS result5, is_inventory AS result6, is_otherdocuments AS result7 FROM user WHERE username = ?", data).Scan(&result, &result2, &result3, &result4, &result5, &result6, &result7)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
        "result2": result2,
        "result3": result3,
        "result4": result4,
        "result5": result5,
        "result6": result6,
        "result7": result7,
			})

    } else if method == "GetMaxEntrycode" {
      err = db.QueryRow("SELECT MAX(entryCodeNo) AS result FROM incoming WHERE substring(entrycodeNo, 1, 2) = ?", data).Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
    }
  })


  router.POST("/api/SaveAccount/:data/:data2/:data3/:data4/:data5/:data6/:data7/:data8", func(c *gin.Context) {
    data := c.Param("data")
    data2 := c.Param("data2")
    data3 := c.Param("data3")
    data4 := c.Param("data4")
    data5 := c.Param("data5")
    data6 := c.Param("data6")
    data7 := c.Param("data7")
    data8 := c.Param("data8")

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO user (userid, employeeName, username, password, is_incoming, is_outgoing, is_releasing, is_inventory, is_otherdocuments, is_monitoring, is_admin) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, 0, 0)")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(data, data2, data3, data4, data5, data6, data7, data8)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Registering")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Registering")
    }
  })



  // router.POST("/api/SaveIncoming/:data/:data2/:data3/:data4/:data5/:data6/:data7/:data8", func(c *gin.Context) {
  //   data := c.Param("data")
  //   data2 := c.Param("data2")
  //   data3 := c.Param("data3")
  //   data4 := c.Param("data4")
  //   data5 := c.Param("data5")
  //   data6 := c.Param("data6")
  //   data7 := c.Param("data7")
  //   data8 := c.Param("data8")

  //   c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
  //   c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
  //   c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
  //   c.Writer.Header().Set("X-Frame-Options", "DENY")
  //   c.Writer.Header().Set("X-Download-Options", "noopen")
  //   c.Writer.Header().Set("Referrer-Policy", "no-referrer")

  //   dbpost, err := db.Prepare("INSERT INTO incoming (entryCodeNo, receivedDate, comType, sourceName, subjectInfo, subDetails, attachments, bo_notes, tag_filing, page_no, folder_no) VALUES (?, ?, '', @name, @subinfo, @subdetails, @attachment, @notes, @tag, @page, @folder)")
  //   if err != nil {
  //     panic(err.Error())
  //   }
  //   defer dbpost.Close()

  //   exec, err := dbpost.Exec(data, data2, data3, data4, data5, data6, data7, data8)
  //   if err != nil {
  //     panic(err.Error())
  //   }

  //   affect, err := exec.RowsAffected()
  //   if err != nil {
  //     panic(err.Error())
  //   }

  //   if affect > 0 {
  //     c.String(http.StatusOK, "Success on Registering")
  //   } else {
  //     c.String(http.StatusInternalServerError, "Failed on Registering")
  //   }
  // })

  router.Run(":8081")
}
