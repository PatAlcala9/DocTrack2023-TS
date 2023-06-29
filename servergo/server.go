package main

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// DEV
var connection string = "root:superuser@tcp(localhost:3306)/ocbodoctracksys"

// SERVER
// var connection string = "iips:iipsuser@tcp(192.168.7.100:3306)/ocbodoctracksys"

func main() {
	connect()
}

func connect() {
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

  // gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
  router.Use(cors.Default())
  // router.LoadHTMLGlob("templates/**")

  // router.GET("/", func(c *gin.Context) {
  //   c.HTML(http.StatusOK, "index.html", gin.H{})
  // })

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

    } else if method == "GetFileReceivedDoc" {
      var result2, result3, result4, result5, result6, result7 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}
      array6 := []string{}
      array7 := []string{}

      results, err := db.Query("SELECT IFNULL(folder_no, '') AS result, IFNULL(page_no, '') AS result2, IFNULL(entryCodeNo, '') AS result3, IFNULL(comType, '') AS result4, IFNULL(sourceName, '') AS result5, IFNULL(subjectInfo, '') AS result6, IFNULL(receivedDate, '') AS result7 FROM incoming WHERE folder_no IS NOT NULL AND folder_no <> '' LIMIT 100")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4, &result5, &result6, &result7)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
        array5 = append(array5, result5)
        array6 = append(array6, result6)
        array7 = append(array7, result7)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
        "result5": array5,
        "result6": array6,
        "result7": array7,
			})

    } else if method == "GetIncomingDesc" {
      var result2, result3, result4 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}

      results, err := db.Query("SELECT IFNULL(entryCodeNo, '') AS result, IFNULL(receivedDate, '') AS result2, IFNULL(sourceName, '') AS result3, IFNULL(subjectInfo, '') AS result4 FROM incoming ORDER BY entryCodeNo DESC LIMIT 100")
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

    } else if method == "GetFileOutgoingDoc" {
      var result2, result3, result4, result5, result6, result7, result8 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}
      array6 := []string{}
      array7 := []string{}
      array8 := []string{}

      results, err := db.Query("SELECT IFNULL(folder_no, '') AS result, IFNULL(page_no, '') AS result2, IFNULL(referenceNo, '') AS result3, IFNULL(comType, '') AS result4, IFNULL(respondent, '') AS result5, IFNULL(subjectInfo, '') AS result6, IFNULL(date_released, '') AS result7, IFNULL(date_received, '') AS result8 FROM outgoing WHERE folder_no IS NOT NULL AND folder_no <> '' LIMIT 100")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result, &result2, &result3, &result4, &result5, &result6, &result7, &result8)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
        array2 = append(array2, result2)
        array3 = append(array3, result3)
        array4 = append(array4, result4)
        array5 = append(array5, result5)
        array6 = append(array6, result6)
        array7 = append(array7, result7)
        array8 = append(array8, result8)
			}
			c.JSON(http.StatusOK, gin.H{
				"result": array,
        "result2": array2,
        "result3": array3,
        "result4": array4,
        "result5": array5,
        "result6": array6,
        "result7": array7,
        "result8": array8,
			})

    } else if method == "GetFileAddIncoming" {
      var result2, result3, result4 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}

      results, err := db.Query("SELECT IFNULL(entryCodeNo, '') AS result, IFNULL(sourceName, '') AS result2, IFNULL(subjectInfo, '') AS result3, IFNULL(receivedDate, '') AS result4 FROM incoming WHERE folder_no IS NULL AND page_no = 0 OR folder_no = '' ORDER BY entryCodeNo DESC LIMIT 50")
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

    } else if method == "GetSources" {
      array := []string{}

      results, err := db.Query("SELECT source_desc AS result FROM source_complaint ORDER BY source_desc ASC")
      if err != nil {
        panic(err.Error())
      }

      for results.Next() {
				err = results.Scan(&result)
				if err != nil {
					panic(err.Error())
				}
				array = append(array, result)
			}

			c.JSON(http.StatusOK, gin.H{
				"result": array,
			})

    } else if method == "GetLatestRespondentID" {
      err = db.QueryRow("SELECT MAX(respondent_infoid) FROM respondent_info").Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})

    } else if method == "GetMaxComplaintCode" {
      err = db.QueryRow("SELECT COALESCE(MAX(complaint_code), '') AS result FROM complaint_info").Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})

    } else if method == "GetComplaintList" {
      var result2, result3, result4, result5 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}

			results, err := db.Query("SELECT c.complaint_code AS result, s.source_desc AS result2, c.complaintant_name AS result3, st.status AS result4, st.date_transacted AS result5 FROM complaint_info c, source_complaint s, complaint_status st WHERE c.source_complaintid = s.source_complaintid AND c.complaint_statusid = st.complaint_statusid")
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

    } else if method == "GetAttachmentList" {
      var result2 string
      array := []string{}
      array2 := []string{}

			results, err := db.Query("SELECT UPPER(description) AS result, ref_attachmentid AS result2 FROM ref_attachment ORDER BY description ASC")
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
      var result2, result3, result4, result5, result6, result7, result8 string

      err = db.QueryRow("SELECT userid AS result, employeeName AS result2, is_incoming AS result3, is_outgoing AS result4, is_releasing AS result5, is_inventory AS result6, is_otherdocuments AS result7, is_complaint AS result8 FROM user WHERE username = ?", data).Scan(&result, &result2, &result3, &result4, &result5, &result6, &result7, &result8)
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
        "result8": result8,
			})

    } else if method == "GetMaxEntryCode" {
      err = db.QueryRow("SELECT MAX(entryCodeNo) AS result FROM incoming WHERE SUBSTRING(entrycodeNo, 1, 2) = ?", data).Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})

    } else if method == "SearchIncoming" {
      var result2, result3, result4, result5, result6 string

      err = db.QueryRow("SELECT IFNULL(receivedDate, '') AS result, IFNULL(sourceName, '') AS result2, IFNULL(subjectInfo, '') AS result3, IFNULL(subDetails, '') AS result4, IFNULL(attachments, '') AS result5, IFNULL(bo_notes, '') AS result6 from incoming where entryCodeNo = ?", data).Scan(&result, &result2, &result3, &result4, &result5, &result6)
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
			})

    } else if method == "SearchIncomingAction" {
      var result2 string
      array := []string{}
      array2 := []string{}

      results, err := db.Query("SELECT IFNULL(action_date, '') AS result, IFNULL(actionMade, '') AS result2 FROM action_logs WHERE entryCodeNo = ?", data)
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

    } else if method == "SearchIncomingDocLog" {
      var result2 string
      array := []string{}
      array2 := []string{}

      results, err := db.Query("SELECT IFNULL(date_forwarded, '') AS result, IFNULL(forwardedto_name, '') AS result2 FROM doc_logs WHERE entryCodeNo = ?", data)
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

    } else if method == "GetSourceID" {
      decodedString := strings.Replace(data, "~", "/", -1)
      err = db.QueryRow("SELECT source_complaintid AS result FROM source_complaint WHERE source_desc = ?", decodedString).Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})

    } else if method == "GetComplaintSpecific" {
      var result2, result3, result4, result5, result6, result7, result8, result9, result10, result11 string

      err = db.QueryRow("SELECT s.source_desc AS result, c.complaintant_name AS result2, c.complaintant_contact AS result3, c.locationOfconstruction AS result4, c.date_received AS result5, c.details AS result6, r.respondent_name AS result7, r.respondent_contact AS result8, r.respondent_location AS result9, st.status AS result10, st.date_transacted AS result11 FROM complaint_info c, source_complaint s, respondent_info r, complaint_status st WHERE c.source_complaintid = s.source_complaintid AND c.respondent_infoid = r.respondent_infoid AND c.complaint_code = st.complaint_code AND c.complaint_code = ?", data).Scan(&result, &result2, &result3, &result4, &result5, &result6, &result7, &result8, &result9, &result10, &result11)
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
        "result8": result8,
        "result9": result9,
        "result10": result10,
        "result11": result11,
			})

    } else if method == "GetLatestStatus" {
      err = db.QueryRow("SELECT COALESCE(MAX(complaint_statusid), '') AS result FROM complaint_status WHERE complaint_code = ?", data).Scan(&result)
      if err != nil {
        panic(err.Error())
      }

			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})

    } else if method == "GetComplaintListFiltered" {
      var result2, result3, result4, result5 string
      array := []string{}
      array2 := []string{}
      array3 := []string{}
      array4 := []string{}
      array5 := []string{}

			results, err := db.Query("SELECT c.complaint_code AS result, s.source_desc AS result2, c.complaintant_name AS result3, st.status AS result4, st.date_transacted AS result5 FROM complaint_info c, source_complaint s, complaint_status st WHERE c.source_complaintid = s.source_complaintid AND c.complaint_statusid = st.complaint_statusid AND c.complaint_code = ?", data)
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



  router.POST("/api/PostAccount", func(c *gin.Context) {
    type AccountData struct {
      Data  string `json:"data"`
      Data2 string `json:"data2"`
      Data3 string `json:"data3"`
      Data4 int `json:"data4"`
      Data5 int `json:"data5"`
      Data6 int `json:"data6"`
      Data7 int `json:"data7"`
      Data8 int `json:"data8"`
      Data9 int `json:"data9"`
    }
    var accountData AccountData
    if err := c.ShouldBindJSON(&accountData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO user (userid, employeeName, username, password, is_incoming, is_outgoing, is_releasing, is_inventory, is_otherdocuments, is_monitoring, is_admin, is_complaint) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, 0, 0, ?)")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(accountData.Data, accountData.Data2, accountData.Data3, accountData.Data4, accountData.Data5, accountData.Data6, accountData.Data7, accountData.Data8, accountData.Data9)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Saving Account")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Saving Account")
    }
  })



  router.POST("/api/PostIncoming", func(c *gin.Context) {
    type IncomingData struct {
      Data  string `json:"data"`
      Data2 string `json:"data2"`
      Data3 string `json:"data3"`
      Data4 string `json:"data4"`
      Data5 string `json:"data5"`
      Data6 string `json:"data6"`
    }
    var incomingData IncomingData
    if err := c.ShouldBindJSON(&incomingData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO incoming (entryCodeNo, receivedDate, comType, sourceName, subjectInfo, subDetails, attachments, bo_notes, tag_filing, page_no, folder_no) VALUES (?, ?, '', ?, ?, ?, ?, '', 0, 0, '')")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(incomingData.Data, incomingData.Data2, incomingData.Data3, incomingData.Data4, incomingData.Data5, incomingData.Data6)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Saving Incoming")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Saving Incoming")
    }
  })



  router.POST("/api/PostRespondent", func(c *gin.Context) {
    type RespondentData struct {
      Data  string `json:"data"`
      Data2 string `json:"data2"`
      Data3 string `json:"data3"`
    }
    var respondentData RespondentData
    if err := c.ShouldBindJSON(&respondentData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO respondent_info (respondent_infoid, respondent_name, respondent_contact, respondent_location) VALUES (NULL, ?, ?, ?)")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(respondentData.Data, respondentData.Data2, respondentData.Data3)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Saving Respondent")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Saving Respondent")
    }
  })



  router.POST("/api/PostComplaint", func(c *gin.Context) {
    type ComplaintData struct {
      Data  string `json:"data"`
      Data2 string `json:"data2"`
      Data3 string `json:"data3"`
      Data4 string `json:"data4"`
      Data5 string `json:"data5"`
      Data6 string `json:"data6"`
      Data7 string `json:"data7"`
      Data8 string `json:"data8"`
      Data9 string `json:"data9"`
    }
    var complaintData ComplaintData
    if err := c.ShouldBindJSON(&complaintData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO complaint_info (complaint_infoid, complaint_code, source_complaintid, complaintant_name, complaintant_contact, date_received, locationOfconstruction, details, respondent_infoid, complaint_statusid) VALUES(NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(complaintData.Data, complaintData.Data2, complaintData.Data3, complaintData.Data4, complaintData.Data5, complaintData.Data6, complaintData.Data7, complaintData.Data8, complaintData.Data9)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Saving Complaint")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Saving Complaint")
    }
  })



  router.POST("/api/PostStatus", func(c *gin.Context) {
    type StatusData struct {
      Data  string `json:"data"`
      Data2 string `json:"data2"`
      Data3 string `json:"data3"`
      Data4 string `json:"data4"`
      Data5 string `json:"data5"`
      Data6 string `json:"data6"`
      Data7 string `json:"data7"`
    }
    var statusData StatusData
    if err := c.ShouldBindJSON(&statusData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO complaint_status (complaint_statusid, complaint_code, date_transacted, status, tagcode, tagword, received_by, remarks) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(statusData.Data, statusData.Data2, statusData.Data3, statusData.Data4, statusData.Data5, statusData.Data6, statusData.Data7)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Saving Status")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Saving Status")
    }
  })



  router.POST("/api/PostAttachment", func(c *gin.Context) {
    type AttachmentData struct {
      Data  string `json:"data"`
      Data2 int `json:"data2"`
    }
    var attachmentData AttachmentData
    if err := c.ShouldBindJSON(&attachmentData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO complaint_attachment (complaint_attachmentid, complaint_code, ref_attachmentid) VALUES (NULL, ?, ?)")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(attachmentData.Data, attachmentData.Data2)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Saving Attachments")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Saving Attachments")
    }
  })



  router.POST("/api/PostEditLog", func(c *gin.Context) {
    type EditLogData struct {
      Data  string `json:"data"`
      Data2 string `json:"data2"`
      Data3 string `json:"data3"`
      Data4 string `json:"data4"`
      Data5 string `json:"data5"`
    }
    var editLogData EditLogData
    if err := c.ShouldBindJSON(&editLogData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    dbpost, err := db.Prepare("INSERT INTO edit_logs (edit_logsid, table, column, old_data, new_data, date_edited) VALUES (NULL, ?, ?, ?, ?, ?)")
    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(editLogData.Data, editLogData.Data2, editLogData.Data3, editLogData.Data4, editLogData.Data5)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, "Success on Saving Edit Log")
    } else {
      c.String(http.StatusInternalServerError, "Failed on Saving Edit Log")
    }
  })



  router.POST("/api/PostUpdate", func(c *gin.Context) {
    type EditData struct {
      Data  string `json:"data"`
      Data2 string `json:"data2"`
      Data3 string `json:"data3"`
    }
    var editData EditData
    if err := c.ShouldBindJSON(&editData); err != nil {
      c.String(http.StatusBadRequest, "Invalid request body")
      return
    }

    c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
    c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
    c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
    c.Writer.Header().Set("X-Frame-Options", "DENY")
    c.Writer.Header().Set("X-Download-Options", "noopen")
    c.Writer.Header().Set("Referrer-Policy", "no-referrer")

    var dbpost *sql.Stmt
    var successMessage, failureMessage string

    switch editData.Data3 {
      case "name":
          dbpost, err = db.Prepare("UPDATE complaint_info SET complaintant_name = ? WHERE complaint_code = ?")
          successMessage = "Success on Updating Complaintant Name"
          failureMessage = "Failed on Updating Complaintant Name"
      case "contact":
          dbpost, err = db.Prepare("UPDATE complaint_info SET complaintant_contact = ? WHERE complaint_code = ?")
          successMessage = "Success on Updating Complaintant Contact"
          failureMessage = "Failed on Updating Complaintant Contact"
      case "location":
        dbpost, err = db.Prepare("UPDATE complaint_info SET locationOfconstruction = ? WHERE complaint_code = ?")
        successMessage = "Success on Updating Complaintant Location"
        failureMessage = "Failed on Updating Complaintant Location"
      case "details":
        dbpost, err = db.Prepare("UPDATE complaint_info SET details = ? WHERE complaint_code = ?")
        successMessage = "Success on Updating Details"
        failureMessage = "Failed on Updating Details"
      case "respondent-name":
        dbpost, err = db.Prepare("UPDATE respondent_info SET respondent_name = ? WHERE respondent_infoid = (SELECT respondent_infoid FROM complaint_info WHERE complaint_code = ? LIMIT 1)")
        successMessage = "Success on Updating Respondent Name"
        failureMessage = "Failed on Updating Respondent Name"
      case "respondent-contact":
        dbpost, err = db.Prepare("UPDATE respondent_info SET respondent_contact = ? WHERE respondent_infoid = (SELECT respondent_infoid FROM complaint_info WHERE complaint_code = ? LIMIT 1)")
        successMessage = "Success on Updating Respondent Contact"
        failureMessage = "Failed on Updating Respondent Contact"
      case "respondent-location":
        dbpost, err = db.Prepare("UPDATE respondent_info SET respondent_location = ? WHERE respondent_infoid = (SELECT respondent_infoid FROM complaint_info WHERE complaint_code = ? LIMIT 1)")
        successMessage = "Success on Updating Respondent Location"
        failureMessage = "Failed on Updating Respondent Location"
      default:
          c.String(http.StatusBadRequest, "Invalid update type")
          return
    }

    if err != nil {
      panic(err.Error())
    }
    defer dbpost.Close()

    exec, err := dbpost.Exec(editData.Data, editData.Data2)
    if err != nil {
      panic(err.Error())
    }

    affect, err := exec.RowsAffected()
    if err != nil {
      panic(err.Error())
    }

    if affect > 0 {
      c.String(http.StatusOK, successMessage)
    } else {
      c.String(http.StatusInternalServerError, failureMessage)
    }
  })



  // router.POST("/api/PostUpdateContact", func(c *gin.Context) {
  //   type EditData struct {
  //     Data  string `json:"data"`
  //     Data2 string `json:"data2"`
  //   }
  //   var editData EditData
  //   if err := c.ShouldBindJSON(&editData); err != nil {
  //     c.String(http.StatusBadRequest, "Invalid request body")
  //     return
  //   }

  //   c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
  //   c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
  //   c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
  //   c.Writer.Header().Set("X-Frame-Options", "DENY")
  //   c.Writer.Header().Set("X-Download-Options", "noopen")
  //   c.Writer.Header().Set("Referrer-Policy", "no-referrer")

  //   dbpost, err := db.Prepare("UPDATE complaint_info SET complaintant_contact = ? WHERE complaint_code = ?")
  //   if err != nil {
  //     panic(err.Error())
  //   }
  //   defer dbpost.Close()

  //   exec, err := dbpost.Exec(editData.Data, editData.Data2)
  //   if err != nil {
  //     panic(err.Error())
  //   }

  //   affect, err := exec.RowsAffected()
  //   if err != nil {
  //     panic(err.Error())
  //   }

  //   if affect > 0 {
  //     c.String(http.StatusOK, "Success on Updating Complaintant Contact")
  //   } else {
  //     c.String(http.StatusInternalServerError, "Failed on Updating Complaintant Contact")
  //   }
  // })


  router.Run(":8081")
}
