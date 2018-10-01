    package main

    import(
      "fmt"

      "database/sql"
      _ "github.com/go-sql-driver/mysql"
      _ "github.com/nfnt/resize"

      "log"
      "time"
      "image"
      "image/png"
      "os"
      "encoding/json"
      "net/http"
      "io/ioutil"
      "bytes"
      "mime/multipart"
       "path/filepath"
       "io"
       "reflect"
       "sync"
       "strconv"
       "html/template"
       //_ "gonum.org/v1/gonum/floats"

        "image/jpeg"


        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/aws/credentials"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/aws/awsutil"
    )


    //func ConnectToDB()
    var wg sync.WaitGroup
    func init() {
            // damn important or else At(), Bounds() functions will
            // caused memory pointer error!!
            image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
    }

    type APIResponse struct {
      bv bool`json:"executionTime"`

    }




    func gogogogo(uid string,imgpath string) {
      // for each row, scan the result into our tag composite object

      defer wg.Done()


                     jsonData := map[string]string{
                       "user_id": uid,
                     }
                    //jsonValue,_ := json.Marshal(jsonData)
                    //fmt.Println(jsonData[user_id])
                    url := "http://staging.vishwamcorp.com/v2/me/reference_ios/"

                      //hit the api with http post request method


                      request, err := newfileUploadRequest(url,jsonData, "image", imgpath)
                      if err != nil {
                        log.Fatal(err)
                      }

                      client := &http.Client{}
                      request.Header.Add("Content-Type", "multipart/form-data")
                      start := time.Now()
                      response_ref_upload, err := client.Do(request)
                      t := time.Now()
                      //response, err := http.Post("http://staging.vishwamcorp.com/v1/me/reference_ios/","application/json", bytes.NewBuffer(jsonValue))

                      elapsed := t.Sub(start)
                        // take the response
                      var resptime float64 = elapsed.Seconds()
                      respCode:=response_ref_upload.StatusCode
                      if err != nil {
                        log.Fatal(err)
                        } else {
                            body := &bytes.Buffer{}
                            _, err := body.ReadFrom(response_ref_upload.Body)
                            if err != nil {
                              log.Fatal(err)
                            }
                              response_ref_upload.Body.Close()//response, err := http.Post("http://staging.vishwamcorp.com/v1/me/reference_ios/","application/json", bytes.NewBuffer(jsonValue))

                              // check the response and Print

                              fmt.Println(response_ref_upload.StatusCode)

                              if response_ref_upload.StatusCode == 200{
                                fmt.Println(uid, "is successfully registered")
                              }
                              //var resCode int = response_ref_upload.StatusCode
                              var status string = "success"
                              //push response data to database User,ref_img_path
                              t:= time.Now()
                              t.Format(time.RFC3339)

                              //db2, _:= sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
                                //db2_local, err := sql.Open("mysql", "root:1666@tcp(127.0.0.1:3306)/Nikhil")

                                db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")

                              /*res, err := db2_local.Query("INSERT INTO MultiReqep1(Uid,refpath,Date,resptime,status,respcode) values(?,?,?,?,?,?);", uid,imgpath,t,resptime,status,respCode)
                                  if err != nil {
                                    panic(err.Error()) // proper error handling instead of panic in your app
                                  }

                                    fmt.Println(res)*/

                                    rescloud, err := db.Query("INSERT INTO MultiReqep1(Uid,refpath,Date,resptime,status,respcode) values(?,?,?,?,?,?);", uid,imgpath,t,resptime,status,respCode)
                                        if err != nil {
                                          panic(err.Error()) // proper error handling instead of panic in your app
                                        }

                                          fmt.Println(rescloud)





                              //fmt.Println(response_ref_upload.Header)
                              fmt.Println(body)
                            /*body_values,err := json.Marshal(body)
                            if err!= nil {
                              panic(err)
                            }

                              fmt.Println(body_values)*/
                              //fmt.Println("variable type of time:",reflect.TypeOf(elapsed))
                              //fmt.Println(elapsed)

                  }

    }






    func get(body []byte) (*APIResponse, error) {
        var ss = new(APIResponse)
        err := json.Unmarshal(body, &ss.bv)
        if(err != nil){
            fmt.Println("whoops:", err)
        }
        return ss, err
    }


    type bs struct {
        referenceBucket   string
        referenceKey string
    }

    func gestureUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {

      var u1 string = params["user_id"]
      fmt.Println(u1)
        var local_image_path1 string = "./gesimages/"+u1+".jpg"

        imgimg1, _ := os.Create(local_image_path1)
        defer imgimg1.Close()

        respimgaws1, _ := http.Get(path)
        defer respimgaws1.Body.Close()
        //fmt.Println("AWS IMAGE:",respimgaws1.Body)
        /*b, err := io.Copy(imgimg1, respimgaws1.Body)
        if err != nil {
           fmt.Println(err)
        }
        fmt.Println("File size: ", b)

       file, err := os.Open(local_image_path1)
       if err != nil {
          return nil, err
       }
       defer file.Close()*/

       body := &bytes.Buffer{}
       writer := multipart.NewWriter(body)
       part, err := writer.CreateFormFile(paramName, filepath.Base(path))
       if err != nil {
          return nil, err
       }
       _, err = io.Copy(part, respimgaws1.Body)

       for key, val := range params {
          _ = writer.WriteField(key, val)
       }

       err = writer.Close()
       if err != nil {
          return nil, err
       }

       req, err := http.NewRequest("POST", uri, body)
       req.Header.Set("Content-Type", writer.FormDataContentType())
       return req, err
    }

    func newfileUploadRequest(url string, params map[string]string, paramName, path string) (*http.Request, error) {


      var u string = params["user_id"]
        var local_image_path string = "./images/"+u+".jpg"

      imgimg, _ := os.Create(local_image_path)
      defer imgimg.Close()

      respimgaws, _ := http.Get(path)
      defer respimgaws.Body.Close()
      //fmt.Println("AWS IMAGE:",respimgaws)
      b, err := io.Copy(imgimg, respimgaws.Body)
      if err != nil {
         fmt.Println(err)
      }

      fmt.Println("File size: ", b)
       file, err := os.Open(local_image_path)
       if err != nil {
          return nil, err
       }
       defer file.Close()

       body := &bytes.Buffer{}
       writer := multipart.NewWriter(body)
       part, err := writer.CreateFormFile(paramName, filepath.Base(local_image_path))
       if err != nil {
          return nil, err
       }
       _, err = io.Copy(part, file)

       for key, val := range params {
          _ = writer.WriteField(key, val)
       }

       err = writer.Close()
       if err != nil {
          return nil, err
       }

       req, err := http.NewRequest("POST", url, body)
       req.Header.Set("Content-Type", writer.FormDataContentType())
       return req, err
    }

    func newMultiFileUploadRequest(uri string, params map[string]string, paramName, path string,paramName2, path2 string) (*http.Request, error) {



     // think about randon user id here instead of taking in the parameters
      var u1 string = params["user_id"]
      fmt.Println(u1)
        var local_image_path1 string = "./images/ep4/"+u1+".jpg"

      imgimg1, _ := os.Create(local_image_path1)
      defer imgimg1.Close()

      respimgaws1, _ := http.Get(path)
      defer respimgaws1.Body.Close()
      //fmt.Println("AWS IMAGE:",respimgaws1.Body)
      b, err := io.Copy(imgimg1, respimgaws1.Body)
      if err != nil {
         fmt.Println(err)
      }




      var u2 string = params["username"]
      fmt.Println(u2)
        var local_image_path2 string = "./images/"+u2+".jpg"

      imgimg2, _ := os.Create(local_image_path2)
      defer imgimg2.Close()

      respimgaws2, _ := http.Get(path2)
      defer respimgaws2.Body.Close()
      //fmt.Println("AWS IMAGE:",respimgaws2.Body)
      b2, err := io.Copy(imgimg2, respimgaws2.Body)
      if err != nil {
         fmt.Println(err)
      }

      fmt.Println("File size: ", b)

      fmt.Println("File size: ", b2)
       file, err := os.Open(local_image_path1)
       if err != nil {
          return nil, err
       }
       defer file.Close()

       file2, err := os.Open(local_image_path2)
       if err != nil {
          return nil, err
       }
       defer file2.Close()

       body := &bytes.Buffer{}
       writer := multipart.NewWriter(body)
       part, err := writer.CreateFormFile(paramName, filepath.Base(local_image_path1))
       if err != nil {
          return nil, err
       }
       _, err = io.Copy(part, file)


       part2, err := writer.CreateFormFile(paramName2, filepath.Base(local_image_path2))
       if err != nil {
          return nil, err
       }
       _, err = io.Copy(part2, file2)




       err = writer.Close()
       if err != nil {
          return nil, err
       }



       req, err := http.NewRequest("POST", uri, body)
       req.Header.Set("Content-Type", writer.FormDataContentType())
       return req, err
    }

    func ep4ReqResp(w http.ResponseWriter,r *http.Request){
      //connect to db

      type Db_values struct
      {
        refpath string
        curimgpath string
        refus string
        curuid string

      }



      db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
        //db_local, err := sql.Open("mysql", "root:1666@tcp(127.0.0.1:3306)/Nikhil")

      if err != nil {
        log.Print(err.Error())
      }

      defer db.Close()

      results, err := db.Query("SELECT usid,username,image_path,pres_img_path FROM fat.regtable,fat.presimgtable WHERE regtable.usid = presimgtable.username;")
      if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
      }


      for results.Next() {


        var expop bool
        var actop bool
        var Values Db_values

        var testresult string
        // for each row, scan the result into our tag composite object
        err = results.Scan(&Values.refus,&Values.curuid,&Values.refpath, &Values.curimgpath)
        if err != nil {
          panic(err.Error()) // proper error handling instead of panic in your app
        }
                    // and then print out the tag's Name attribute
        //var render_data string = uid
    //fmt.Println("Uid: ",Values.Uid,"Image path: ",Values.Img_path)


      extraParams := map[string]string{
        "usid":Values.refus,
        "username":Values.curuid,
         }
      url := "http://staging.vishwamcorp.com/v1/direct_match"

      request, err := newMultiFileUploadRequest(url, extraParams, "image1", Values.refpath, "image2", Values.curimgpath)

      request.Header.Add("Content-Type", "multipart/form-data")
      if err != nil {
        fmt.Println(err)
         log.Fatal(err)

      }
      client := &http.Client{}
      startx := time.Now()
      resp, err := client.Do(request)
      tx := time.Now()
      elapsedx := tx.Sub(startx)
      var resptime_ep4 float64 = elapsedx.Seconds()
      if err != nil {
         log.Fatal(err)
      } else {
       // body := &bytes.Buffer{}
         body_values, err := ioutil.ReadAll(resp.Body)
         if err != nil {
            log.Fatal(err)
         }

         s, err := get([]byte(body_values))
         resp.Body.Close()
         fmt.Println(resp.StatusCode)
         //fmt.Println(resp.Header)

         //_,err := json.Unmarshal(body,&resgot)
         fmt.Println("Cracked",s.bv)
         expop = true
         actop = s.bv
         if actop==expop{
           fmt.Println("yahoo")
           testresult = "pass"

         }else{
           testresult = "fail"
         }
         fmt.Println(resp.Body)

         t:= time.Now()

          t.Format(time.RFC3339)
      /*   res, err := db_local.Query("INSERT INTO EP4_RESP_NEW(Date,exp_op,actual_op,test_result,respcode, resptime) values(?,?,?,?,?,?);",t,expop,actop,testresult,resp.StatusCode,resptime_ep4)
             if err != nil {
               panic(err.Error()) // proper error handling instead of panic in your app
             }

               fmt.Println(res)*/

               rescloud, err := db.Query("INSERT INTO EP4_RESP(DATE,EXPOP,ACTUALOP,TESTRESULT,RESPCODE, RESPTIME) values(?,?,?,?,?,?);",t,expop,actop,testresult,resp.StatusCode,resptime_ep4)
                   if err != nil {
                     panic(err.Error()) // proper error handling instead of panic in your app
                   }

                     fmt.Println(rescloud)





      }
    }
    results2, err := db.Query("SELECT usid,username,image_path,pres_img_path FROM fat.regtable,fat.presimgtable WHERE regtable.usid != presimgtable.username;")
    if err != nil {
      panic(err.Error()) // proper error handling instead of panic in your app
    }


    for results2.Next() {

      var expop2 bool
      var actop2 bool
      var Values2 Db_values

      var testresult2 string
      // for each row, scan the result into our tag composite object
      err = results2.Scan(&Values2.refus,&Values2.curuid,&Values2.refpath, &Values2.curimgpath)
      if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
      }
                  // and then print out the tag's Name attribute
      //var render_data string = uid
  //fmt.Println("Uid: ",Values.Uid,"Image path: ",Values.Img_path)


    extraParam := map[string]string{

      "usid":Values2.refus,
      "username":Values2.curuid,

       }
    url := "http://staging.vishwamcorp.com/v2/direct_match_ios"

    requestreq, err := newMultiFileUploadRequest(url, extraParam, "image1", Values2.refpath, "image2", Values2.curimgpath)

    requestreq.Header.Add("Content-Type", "multipart/form-data")
    if err != nil {
      fmt.Println(err)
       log.Fatal(err)

    }
    clientc := &http.Client{}
    startx := time.Now()
    respresp, err := clientc.Do(requestreq)
    tx := time.Now()
    elapsedx := tx.Sub(startx)
    var resptime_ep4 float64 = elapsedx.Seconds()
    if err != nil {
       log.Fatal(err)
    } else {
     //  body := &bytes.Buffer{}
       body_values2, err := ioutil.ReadAll(respresp.Body)
       if err != nil {
          log.Fatal(err)
       }

       s2, err := get([]byte(body_values2))
       respresp.Body.Close()
       fmt.Println(respresp.StatusCode)
       //fmt.Println(respresp.Header)

       //_,err := json.Unmarshal(body,&resgot)
       fmt.Println("Cracked",s2.bv)
       expop2 = false
       actop2 = s2.bv
       if s2.bv==expop2{
         fmt.Println("yahoo")
         testresult2 = "pass"

       }else{
         testresult2 = "fail"
       }
       fmt.Println(respresp.Body)
       t:= time.Now()
       t.Format(time.RFC3339)
      /*  res, err := db_local.Query("INSERT INTO EP4_RESP_NEW(Date,exp_op,actual_op,test_result,respcode,resptime) values(?,?,?,?,?,?);",t,expop2,actop2,testresult2,respresp.StatusCode,resptime_ep4)
           if err != nil {
             panic(err.Error()) // proper error handling instead of panic in your app
           }

             fmt.Println(res)*/


             rescloud, err := db.Query("INSERT INTO EP4_RESP(DATE,EXPOP,ACTUALOP,TESTRESULT,RESPCODE, RESPTIME) values(?,?,?,?,?,?);",t,expop2,actop2,testresult2,respresp.StatusCode,resptime_ep4)
                 if err != nil {
                   panic(err.Error()) // proper error handling instead of panic in your app
                 }

                   fmt.Println(rescloud)

    }
  }

  fmt.Fprintf(w,"Direct Match Images End point is successfully tested")

  }

    /*func gestureEndPointRequest(url string, params map[string]interface{}, paramName, path string) (*http.Request, error) {
       file, err := os.Open(path)
       if err != nil {
          return nil, err
       }
       defer file.Close()

       body := &bytes.Buffer{}
       writer := multipart.NewWriter(body)
       part, err := writer.CreateFormFile(paramName, filepath.Base(path))
       if err != nil {
          return nil, err
       }
       _, err = io.Copy(part, file)

        //for key, val:= range params {
          //_ = writer.WriteField(key, val)
       //}

    for key, param:= range params {

      value, ok := param.(int)
      //dosomethingWith(value)
      if(ok==true){
        _ = writer.WriteField(key, value)
      } else{
        value,_ := param.(string)
        _ = writer.WriteField(key, value)

      }
  }
      //_ = writer.WriteField(key, val)


       err = writer.Close()
       if err != nil {
          return nil, err
       }

       req, err := http.NewRequest("POST", url, body)
       req.Header.Set("Content-Type", writer.FormDataContentType())
       return req, err
    }*/
    func ep5ReqResp(w http.ResponseWriter,r *http.Request){


      type Db_values struct
      {
        Uid string
        gno int
        gespath string
        expop int

      }

      db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
      //db_local, err := sql.Open("mysql", "root:1666@tcp(127.0.0.1:3306)/Nikhil")

      if err != nil {
        log.Print(err.Error())
      }

      defer db.Close()

      // pass a select query to reteieve uid and image path


      results, err := db.Query("SELECT * FROM gesimgtable;")
      if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
      }

      // parse through the database object returned from databases

      for results.Next() {

      var gesval Db_values
        // for each row, scan the result into our tag composite object
        err = results.Scan(&gesval.Uid,&gesval.gespath,&gesval.gno,&gesval.expop)
        if err != nil {
          panic(err.Error()) // proper error handling instead of panic in your app
        }
                    // and then print out the tag's Name attribute
        //var render_data string = uid
      //fmt.Println("Uid: ",Values.Uid)



      //connect to local database



      //for every row send the request to the api
                      var User string = gesval.Uid
                      var gn string = strconv.Itoa(gesval.gno)
                      var gg int = gesval.expop
                      var  exp string
                      if(gg==1){
                        exp = "ok"
                      }else{
                        exp="fail"
                      }


                      var gesturepath string = gesval.gespath
                      var testresult string
                      var stst string

                      var actop string


       extraParams := map[string]string{

          "gesture": gn,
           "delay":"500",
           "version":"1",
          "user_id": User,
       }
       url := "http://staging.vishwamcorp.com/v2/single_gesture_ios"

       fmt.Println(gesturepath)

       request, err := newfileUploadRequest(url, extraParams, "image", gesturepath)
       //request.Header.Add("Content-Type", "multipart/form-data")
       if err != nil {
          log.Fatal(err)
       }
        client := &http.Client{}
       startx := time.Now()
       resp, err := client.Do(request)
       tx := time.Now()
       elapsedx := tx.Sub(startx)
       var resptime_ep5 float64 = elapsedx.Seconds()
       if err != nil {
          log.Fatal(err)
       } else {
        // body := &bytes.Buffer{}
          body_values_gp, err := ioutil.ReadAll(resp.Body)
          if err != nil {
             log.Fatal(err)
          }



          var ep5_resp map[string]string
          //unmarchal extracts json data into map variable
          json.Unmarshal(body_values_gp, &ep5_resp)


          actop = ep5_resp["status"]

          //fmt.Println("VERY VERY IMPORTANT:",body_values_ep3)
          if exp == actop{
            testresult = "pass"
          } else{
            testresult = "fail"
          }
          resp.Body.Close()
          fmt.Println(resp.StatusCode)
          //fmt.Println(resp.Header)
          //fmt.Println(body_values_ep3)
       }



        if resp.StatusCode == 200{
          //fmt.Println("UserId is successfully retrieved")
          stst = "sucess"
          //resp_code = resp.StatusCode

        }else{
          stst = "failure"
          //resp_code = resp.StatusCode
        }

        t:= time.Now()

         t.Format(time.RFC3339)

       /*   res, err := db_local.Query("INSERT INTO EP4_RESP_NEW(Date,exp_op,actual_op,test_result,respcode, resptime) values(?,?,?,?,?,?);",t,expop,actop,testresult,resp.StatusCode,resptime_ep4)
              if err != nil {
                panic(err.Error()) // proper error handling instead of panic in your app
              }

                fmt.Println(res)*/

                rescloud, err := db.Query("INSERT INTO EP5_RESP(DATE,GNO,GESPATH,EXPRES,ACTRES,TESTRES,RESPCODE,STATUS,RESPTIME) values(?,?,?,?,?,?,?,?,?);",t,gn,gesturepath,exp,actop,testresult,resp.StatusCode,stst,resptime_ep5)
                    if err != nil {
                      panic(err.Error()) // proper error handling instead of panic in your app
                    }

                      fmt.Println(rescloud)

    }

    fmt.Fprintf(w, "gesture upload endpoint is successfully tested")
    }

    func ep1ReqResp(w http.ResponseWriter,r *http.Request){
      type Db_values struct
      {
        Uid string
        Img_path string

      }
       var Values Db_values

      //connect to local database

      db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
      //db_local, err := sql.Open("mysql", "root:1666@tcp(127.0.0.1:3306)/Nikhil")

      if err != nil {
        log.Print(err.Error())
      }

      //defer db_local.Close()

      // pass a select query to reteieve uid and image path


      results, err := db.Query("SELECT * FROM regtable;")
      if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
      }

      // parse through the database object returned from databases

      for results.Next() {


        // for each row, scan the result into our tag composite object
        err = results.Scan(&Values.Uid, &Values.Img_path)
        if err != nil {
          panic(err.Error()) // proper error handling instead of panic in your app
        }
                    // and then print out the tag's Name attribute
        //var render_data string = uid
    fmt.Println("Uid: ",Values.Uid,"Image path: ",Values.Img_path)

    //for every row send the request to the api
                      var User = Values.Uid
                      fmt.Println("hahahahahah:"+User)
                      var ref_img_path string = Values.Img_path
                      fmt.Println("hahahahahah:"+ref_img_path)



                       jsonData := map[string]string{
                         "user_id": User,
                       }
                      //jsonValue,_ := json.Marshal(jsonData)
                      //fmt.Println(jsonData[user_id])
                      url := "http://staging.vishwamcorp.com/v2/me/reference_ios/"

                        //hit the api with http post request method


                        request, err := newfileUploadRequest(url,jsonData, "image", ref_img_path)
                        if err != nil {
                          log.Fatal(err)
                        }

                        client := &http.Client{}
                        request.Header.Add("Content-Type", "multipart/form-data")
                        start := time.Now()
                        response_ref_upload, err := client.Do(request)
                        t := time.Now()
                        //response, err := http.Post("http://staging.vishwamcorp.com/v1/me/reference_ios/","application/json", bytes.NewBuffer(jsonValue))

                        elapsed := t.Sub(start)
                        var resptime_ep1 float64 = elapsed.Seconds()

                          // take the response
                        if err != nil {
                          log.Fatal(err)
                          } else {
                              body := &bytes.Buffer{}
                              _, err := body.ReadFrom(response_ref_upload.Body)
                              if err != nil {
                                log.Fatal(err)
                              }
                                response_ref_upload.Body.Close()//response, err := http.Post("http://staging.vishwamcorp.com/v1/me/reference_ios/","application/json", bytes.NewBuffer(jsonValue))

                                // check the response and Print

                                fmt.Println(response_ref_upload.StatusCode)
                                var status string
                                var resCode int
                                if response_ref_upload.StatusCode == 200{
                                  fmt.Println(User, "is successfully registered")
                                   status = "success"
                                   resCode = response_ref_upload.StatusCode
                                }else{
                                  status="failure"
                                  resCode = response_ref_upload.StatusCode
                                }


                                //push response data to database User,ref_img_path
                                t:= time.Now()

                                 t.Format(time.RFC3339)

                                /*res, err := db_local.Query("INSERT INTO EP1_RESP(uid,refpath,resp_code,status,Date,resptime) values(?,?,?,?,?,?);", User,ref_img_path,resCode,status,t,resptime_ep1)
                                    if err != nil {
                                      panic(err.Error()) // proper error handling instead of panic in your app
                                    }

                                      fmt.Println(res)*/

                                      rescloud, err := db.Query("INSERT INTO EP1_RESP(UDI,DATE,REFPATH,RESP_CODE,STATUS,RESPTIME) values(?,?,?,?,?,?);", User,t,ref_img_path,resCode,status,resptime_ep1)
                                          if err != nil {
                                            panic(err.Error()) // proper error handling instead of panic in your app
                                          }

                                            fmt.Println(rescloud)




                                //fmt.Println(response_ref_upload.Header)
                                fmt.Println(body)
                              /*body_values,err := json.Marshal(body)
                              if err!= nil {
                                panic(err)
                              }

                                fmt.Println(body_values)*/
                                fmt.Println("variable type of time:",reflect.TypeOf(elapsed))

                                fmt.Println(elapsed)

                    }



    }

      fmt.Fprintf(w,"Reference upload image is successfully tested...!!!")


    }

    func ep3ReqResP(w http.ResponseWriter,r *http.Request){

                  type Db_values struct
                  {
                    Uid string
                    image_path string

                  }
                   var Values Db_values

                   var expop string
                   var actop string
                   var testresultep3 string
                   var status string
                    var resp_code int
                  //connect to local database

                  db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
                  //db_local, err := sql.Open("mysql", "root:1666@tcp(127.0.0.1:3306)/Nikhil")

                  if err != nil {
                    log.Print(err.Error())
                  }

                  defer db.Close()

                  // pass a select query to reteieve uid and image path


                  results, err := db.Query("SELECT * FROM regtable;")
                  if err != nil {
                    panic(err.Error()) // proper error handling instead of panic in your app
                  }

                  // parse through the database object returned from databases

                  for results.Next() {

                    // for each row, scan the result into our tag composite object
                    err = results.Scan(&Values.Uid,&Values.image_path)
                    if err != nil {
                      panic(err.Error()) // proper error handling instead of panic in your app
                    }
                                // and then print out the tag's Name attribute
                    //var render_data string = uid



                  //connect to local database



                //for every row send the request to the api
                                  var User = Values.Uid


                                  fmt.Println("Sending request to EP 3")
                                  extraParams := map[string]string{
                                    "user_id": User,

                                     }
                                  url := "http://staging.vishwamcorp.com/v2/face_lookup_ios"

                                  request, err := newfileUploadRequest(url, extraParams, "image", Values.image_path)

                                  request.Header.Add("Content-Type", "multipart/form-data")
                                  if err != nil {
                                     log.Fatal(err)
                                  }
                                  client := &http.Client{}
                                    startx := time.Now()
                                  resp, err := client.Do(request)


                                  tx := time.Now()
                                  elapsedx := tx.Sub(startx)
                                  var resptime_ep3 float64 = elapsedx.Seconds()
                                  fmt.Println("The Response Time for EP3:",elapsedx)

                                  if err != nil {
                                     log.Fatal(err)
                                  } else {
                                    body_values_ep3, err := ioutil.ReadAll(resp.Body)
                                      //body := &bytes.Buffer{}
                                    //_, err := body.ReadFrom(resp.Body)
                                     //var m map[string]string

                                     //sep3, err := getep3([]byte(body_values_ep3))
                                     //mmm,err:= json.Marshal(body_values_ep3)
                                     if err != nil {
                                        log.Fatal(err)
                                     }

                                     //creating maps to store JSON data

                                     var ep3_resp map[string]string
                                     //unmarchal extracts json data into map variable
                                     json.Unmarshal(body_values_ep3, &ep3_resp)
                                     fmt.Println(ep3_resp["userId"])
                                     expop = User
                                     actop = ep3_resp["userId"]

                                     //fmt.Println("VERY VERY IMPORTANT:",body_values_ep3)
                                     if expop == actop{
                                       testresultep3 = "pass"
                                     } else{
                                       testresultep3 = "fail"
                                     }
                                     resp.Body.Close()
                                     fmt.Println(resp.StatusCode)
                                     //fmt.Println(resp.Header)
                                     //fmt.Println(body_values_ep3)
                                  }


                                   if resp.StatusCode == 200{
                                     fmt.Println("UserId is successfully retrieved")
                                     status = "sucess"
                                     resp_code = resp.StatusCode

                                   }else{
                                     status = "failure"
                                     resp_code = resp.StatusCode
                                   }
                                   //fmt.Println(status,resp_code)
                                   t:= time.Now()

                                    t.Format(time.RFC3339)
                                    /*res, err := db_local.Query("INSERT INTO EP3_RESP(resp_code,status,reptime,Date,expop,actop,testresult) values(?,?,?,?,?,?,?);",resp_code,status,resptime_ep3,t,expop,actop,testresultep3)
                                       if err != nil {
                                         panic(err.Error()) // proper error handling instead of panic in your app
                                       }

                                         fmt.Println(res)*/

                                         rescloud, err := db.Query("INSERT INTO EP3_RESP(DATE,EXPID,ACTID,RESPCODE,RESPTIME,TESTRESULT,STATUS) values(?,?,?,?,?,?,?);",t,expop,actop,resp_code,resptime_ep3,testresultep3,status)
                                            if err != nil {
                                              panic(err.Error()) // proper error handling instead of panic in your app
                                            }

                                              fmt.Println(rescloud)


                }
                fmt.Fprintf(w,"Face look up End point is succesfully tested...!!!!!")


    }

    // Sending Multiple Requests at a time for EP2


    func ep2MultiReq(w http.ResponseWriter,r *http.Request){

      fmt.Fprintf(w,"Hitting Ep2 Multiple Requests at a time")
       type Db_values struct
                {
                  Uid string
                   }
                 var Values Db_values

                 db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat");


              if err != nil {
                 log.Print(err.Error())
                }

                defer db.Close()

                // pass a select query to reteieve uid and image path


                results, err := db.Query("SELECT usid FROM regtable;")
                if err != nil {
                  panic(err.Error()) // proper error handling instead of panic in your app
                }
                for results.Next() {

                        // for each row, scan the result into our tag composite object
                        err = results.Scan(&Values.Uid)
                        if err != nil {
                          panic(err.Error()) // proper error handling instead of panic in your app
                        }

                            var User = Values.Uid
                            wg.Add(1)
                            go ep2GoReq(User)
                            wg.Wait()
                     }



                     type multires struct{
                       nr int
                       avgrt float64
                     }


                     //send the analytics of Multiple Request

                     t := time.Now()
                     a:= strconv.Itoa(t.Year())
                     b:= strconv.Itoa(int(t.Month()))
                     c:= strconv.Itoa(t.Day())
                     d:= a+"-"+b+"-"+c


                    //fmt.Println(d.Format(time.RFC3339))
                     resul, err := db.Query("SELECT Count(*),Max(resptime) FROM MultiReqEP2 WHERE Date = ?;",d)
                     if err != nil {
                       panic(err.Error()) // proper error handling instead of panic in your app
                     }

                     var MmmRrrr multires
                     for resul.Next(){

                       err = resul.Scan(&MmmRrrr.nr,&MmmRrrr.avgrt)
                       fmt.Println()
                       if err != nil {
                         panic(err.Error()) // proper error handling instead of panic in your app
                       }

                     }
                     resu, err := db.Query("INSERT INTO EP2M_RESP(DATE,NO_OF_RESP,TIME) values(?,?,?);",d,MmmRrrr.nr,MmmRrrr.avgrt)
                     if err != nil {
                       panic(err.Error()) // proper error handling instead of panic in your app
                     }

                     fmt.Println(resu)

                       fmt.Fprintf(w,"Done with hitting multiple req-EP2")

    }
    func ep2GoReq(User1 string){
      defer wg.Done()


                                fmt.Println("Sending multiple request to EP 2")
                                url := "http://staging.vishwamcorp.com/v2/me/reference_ios/"
                                url += User1
                                startM := time.Now()
                                responseM, err := http.Get(url)
                                tM := time.Now()

                                defer responseM.Body.Close()
                                elapsedM := tM.Sub(startM)
                                var resptime_ep2M float64 = elapsedM.Seconds()
                                 if err != nil{
                                   fmt.Println("The Http request is failed  with error %s", err)

                                 } else{


                                    fmt.Println("Response code of EP2:")
                                    fmt.Println(responseM.StatusCode)
                                    newImgUrl:= "./images/succes"
                                    newImgUrl += User1+".png"
                                    file, err := os.Create(newImgUrl)
                                    if err != nil {
                                        log.Fatal(err)
                                    }
                                    // Use io.Copy to just dump the response body to the file. This supports huge files
                                    _, err = io.Copy(file, responseM.Body)
                                    if err != nil {
                                        log.Fatal(err)
                                    }
                                    file.Close()
                                    fmt.Println("Success")

                                    fmt.Println("Response time of EP2:",elapsedM)
                                 }

                                 var status string
                                  var resp_code int
                                 if responseM.StatusCode == 200{
                                   fmt.Println("Image is successfully retrieved")
                                   status = "sucess"
                                   resp_code = responseM.StatusCode

                                 }else{
                                   status = "fail"
                                   resp_code = responseM.StatusCode
                                 }



                                  tM.Format(time.RFC3339)

                                db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")

                                     rescloud1, err := db.Query("INSERT INTO MultiReqEP2(UID,DATE,RESPCODE,RESPTIME,STATUS) values(?,?,?,?,?);", User1,tM,resp_code,resptime_ep2M,status)
                                            if err != nil {
                                              panic(err.Error()) // proper error handling instead of panic in your app
                                            }

                                              fmt.Println(rescloud1)

     }

    func ep1MultipleReqResp(w http.ResponseWriter,r *http.Request){
      type Db_values struct
      {
        Uid string
        Img_path string

      }
       var Values Db_values


       type multires struct{
         nr int
         avgrt float64
       }
      //connect to local database


        //db_local, err := sql.Open("mysql", "root:1666@tcp(127.0.0.1:3306)/Nikhil")

        db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat");

      if err != nil {
        log.Print(err.Error())
      }

      //defer db_local.Close()

      // pass a select query to reteieve uid and image path


      results, err := db.Query("SELECT * FROM regtable;")
      if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
      }

      // parse through the database object returned from databases

      for results.Next()  {

        err = results.Scan(&Values.Uid, &Values.Img_path)
        if err != nil {
          panic(err.Error()) // proper error handling instead of panic in your app
        }
                    // and then print out the tag's Name attribute
        //var render_data string = uid
  //fmt.Println("Uid: ",Values.Uid,"Image path: ",Values.Img_path)

    //for every row send the request to the api
                      var User string= Values.Uid
                      //fmt.Println("hahahahahah:"+User)
                      var ref_img_path string = Values.Img_path
                      //fmt.Println("hahahahahah:"+ref_img_path)*/

                      wg.Add(1)

        go gogogogo(User,ref_img_path)


        wg.Wait()


    }

    //send the analytics of Multiple Request

    t := time.Now()
    a:= strconv.Itoa(t.Year())
    b:= strconv.Itoa(int(t.Month()))
    c:= strconv.Itoa(t.Day())
    d:= a+"-"+b+"-"+c


   //fmt.Println(d.Format(time.RFC3339))
    resul, err := db.Query("SELECT Count(*),Max(resptime) FROM MultiReqep1 WHERE Date = ?;",d)
    if err != nil {
      panic(err.Error()) // proper error handling instead of panic in your app
    }

    var MmmRrrr multires
    for resul.Next(){

      err = resul.Scan(&MmmRrrr.nr,&MmmRrrr.avgrt)
      fmt.Println()
      if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
      }

    }
    resu, err := db.Query("INSERT INTO EP1M_RESP(DATE,NO_OF_RESP,TIME) values(?,?,?);",d,MmmRrrr.nr,MmmRrrr.avgrt)
    if err != nil {
      panic(err.Error()) // proper error handling instead of panic in your app
    }

    fmt.Println(resu)


    fmt.Fprintf(w,"Reference_Upload End point is succesfully tested by sending multiple images at a time..!!!!")
  }


    func ep2ReqResp(w http.ResponseWriter,r *http.Request){


            type Db_values struct
            {
              Uid string


            }
             var Values Db_values

            //connect to local database

            db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
          //  db_local, err := sql.Open("mysql", "root:1666@tcp(127.0.0.1:3306)/Nikhil")

            if err != nil {
              log.Print(err.Error())
            }

            defer db.Close()

            // pass a select query to reteieve uid and image path


            results, err := db.Query("SELECT usid FROM regtable;")
            if err != nil {
              panic(err.Error()) // proper error handling instead of panic in your app
            }

            // parse through the database object returned from databases

            for results.Next() {


              // for each row, scan the result into our tag composite object
              err = results.Scan(&Values.Uid)
              if err != nil {
                panic(err.Error()) // proper error handling instead of panic in your app
              }
                          // and then print out the tag's Name attribute
              //var render_data string = uid
          fmt.Println("Uid: ",Values.Uid)



            //connect to local database



          //for every row send the request to the api
                            var User = Values.Uid
                            fmt.Println("hahahahahah:"+User)

                            fmt.Println("Sending request to EP 2")
                            url := "http://staging.vishwamcorp.com/v2/me/reference_ios/"
                            url += User
                            start2 := time.Now()
                            response2, err := http.Get(url)
                            t2 := time.Now()

                            defer response2.Body.Close()
                            elapsed2 := t2.Sub(start2)
                            var resptime_ep2 float64 = elapsed2.Seconds()
                             if err != nil{
                               fmt.Println("The Http request is failed  with error %s", err)

                             } else{
                                //data,_ := ioutil.ReadAll(response2.Body)
                                //fmt.Println(string(data))

                              //  img, _, err := image.Decode(response2.Body)




                                // is this image opaque
                                // op := canvas.Opaque()


                                fmt.Println("Response code of EP2:")
                                fmt.Println(response2.StatusCode)
                                newImgUrl:= "./images/succes"
                                newImgUrl += User+".png"
                                file, err := os.Create(newImgUrl)
                                if err != nil {
                                    log.Fatal(err)
                                }
                                // Use io.Copy to just dump the response body to the file. This supports huge files
                                _, err = io.Copy(file, response2.Body)
                                if err != nil {
                                    log.Fatal(err)
                                }
                                file.Close()
                                fmt.Println("Success")

                                fmt.Println("Response time of EP2:",elapsed2)
                             }

                             var status string
                              var resp_code int
                             if response2.StatusCode == 200{
                               fmt.Println("Image is successfully retrieved")
                               status = "sucess"
                               resp_code = response2.StatusCode

                             }else{
                               status = "fail"
                               resp_code = response2.StatusCode
                             }

                             t := time.Now()
                              t.Format(time.RFC3339)
                             /*res, err := db_local.Query("INSERT INTO EP2_RESP(uid,resp_code,status,resptime,Date) values(?,?,?,?,?);", User,resp_code,status,resptime_ep2,t)
                                 if err != nil {
                                   panic(err.Error()) // proper error handling instead of panic in your app
                                 }

                                   fmt.Println(res)*/


                                   rescloud, err := db.Query("INSERT INTO EP2_RESP(UID,DATE,RESP_CODE,RESPTIME,STATUS) values(?,?,?,?,?);", User,t,resp_code,resptime_ep2,status)
                                       if err != nil {
                                         panic(err.Error()) // proper error handling instead of panic in your app
                                       }

                                         fmt.Println(rescloud)
                                         /*
                                         type mrv struct{
                                           nr int
                                           maxrt float64
                                         }

                                         resul, err := db.Query("SELECT Count(*),Max(resptime) FROM MultiReqep1 WHERE Date = ?;",d)
                                         if err != nil {
                                           panic(err.Error()) // proper error handling instead of panic in your app
                                         }

                                         var MmmRrrr mrv
                                         for resul.Next(){
                                           err = resul.Scan(&MmmRrrr.nr, &MmmRrrr.maxrt)
                                           if err != nil {
                                             panic(err.Error()) // proper error handling instead of panic in your app
                                           }

                                         }

                                         resu, err := db.Query("INSERT INTO mtest(DATE,N0_OF_REQUEST,MAXTIME) values(?,?,?);",d,MmmRrrr.nr,MmmRrrr.maxrt)
                                         if err != nil {
                                           panic(err.Error()) // proper error handling instead of panic in your app
                                         }

                                         fmt.Println(resu)*/







          }


          fmt.Fprintf(w,"Retrieve Reference Image is successfully tested...!!!")

    }




    func ep4MultiReq(w http.ResponseWriter,r *http.Request){

      fmt.Fprintf(w,"ep4MultiReq running")
      type Db_values struct
          {
            refpath string
            curimgpath string
            u1 string
            u2 string

          }

          db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat");
          if err != nil {
            log.Print(err.Error())
          }

          defer db.Close()

          results, err := db.Query("SELECT usid,username,image_path,pres_img_path FROM fat.regtable,fat.presimgtable WHERE regtable.usid = presimgtable.username;")
          if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
          }


          for results.Next() {

            var Values Db_values
            // for each row, scan the result into our tag composite object
           err = results.Scan(&Values.u1, &Values.u2, &Values.refpath, &Values.curimgpath)
            if err != nil {
              panic(err.Error()) // proper error handling instead of panic in your app
            }
                var imgT1=Values.refpath
                 var imgT2=Values.curimgpath
                 wg.Add(1)
                 go ep4GoReqT(imgT1,imgT2,Values.u1,Values.u2)
                 wg.Wait()
            }



          results2, err := db.Query("SELECT usid,username,image_path,pres_img_path FROM fat.regtable,fat.presimgtable WHERE regtable.usid != presimgtable.username;")
          if err != nil {
          panic(err.Error()) // proper error handling instead of panic in your app
        }


        for results2.Next() {
          var Values2 Db_values
          // for each row, scan the result into our tag composite object
          err = results2.Scan(&Values2.u1, &Values2.u2, &Values2.refpath, &Values2.curimgpath)
          if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
          }

            var imgF1=Values2.refpath
             var imgF2=Values2.curimgpath
             wg.Add(1)
             go ep4GoReqT(imgF1,imgF2,Values2.u1,Values2.u2)
              wg.Wait()
             }


             type multires struct{
               nr int
               avgrt float64
             }


             //send the analytics of Multiple Request

             t := time.Now()
             a:= strconv.Itoa(t.Year())
             b:= strconv.Itoa(int(t.Month()))
             c:= strconv.Itoa(t.Day())
             d:= a+"-"+b+"-"+c


            //fmt.Println(d.Format(time.RFC3339))
             resul, err := db.Query("SELECT Count(*),Max(resptime) FROM MultiReqEP4 WHERE Date = ?;",d)
             if err != nil {
               panic(err.Error()) // proper error handling instead of panic in your app
             }

             var MmmRrrr multires
             for resul.Next(){

               err = resul.Scan(&MmmRrrr.nr,&MmmRrrr.avgrt)
               fmt.Println()
               if err != nil {
                 panic(err.Error()) // proper error handling instead of panic in your app
               }

             }
             resu, err := db.Query("INSERT INTO EP4M_RESP(DATE,NO_OF_RESP,TIME) values(?,?,?);",d,MmmRrrr.nr,MmmRrrr.avgrt)
             if err != nil {
               panic(err.Error()) // proper error handling instead of panic in your app
             }

             fmt.Println(resu)

             fmt.Fprintf(w,"done with Ep4 multiple req")


    }

    func ep4GoReqT(ep4imgpathT1 string,ep4imgpathT2 string,u1 string,u2 string){
      defer wg.Done()
      var expop1 bool
       var actop1 bool
       extraParams1 := map[string]string{

         "user_id": u1,
         "username":u2,
             }
          url := "http://staging.vishwamcorp.com/v2/direct_match_ios"
          request, err := newMultiFileUploadRequest(url, extraParams1, "image1", ep4imgpathT1, "image2",ep4imgpathT2)

          //request.Header.Add("Content-Type", "multipart/form-data")
          if err != nil {
            fmt.Println(err)
             log.Fatal(err)
          }
          client := &http.Client{}
          startA := time.Now()
          resp, err := client.Do(request)
          tA := time.Now()
          elapsedA := tA.Sub(startA)
          var resptime_ep4M float64 = elapsedA.Seconds()
          if err != nil {
             log.Fatal(err)
          } else {
           //  body := &bytes.Buffer{}
             body_values, err := ioutil.ReadAll(resp.Body)
             if err != nil {
                log.Fatal(err)
             }
             s, err := get([]byte(body_values))
             resp.Body.Close()
             fmt.Println(resp.StatusCode)
             fmt.Println(resp.Header)
             var testresult string
             //_,err := json.Unmarshal(body,&resgot)
             fmt.Println("Cracked",s.bv)
             expop1 = true
             actop1 = s.bv
             if actop1==expop1{
               fmt.Println("yahoo")
               testresult = "pass"
             }
        // else{
          //     testresult = "fail"
            // }
             fmt.Println(resp.Body)

             tA:= time.Now()

              tA.Format(time.RFC3339)
              db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat");
             res, err := db.Query("INSERT INTO MultiReqEP4(DATE,RESPCODE,RESPTIME,STATUS) values(?,?,?,?);",tA,resp.StatusCode,resptime_ep4M,testresult)
                 if err != nil {
                   panic(err.Error()) // proper error handling instead of panic in your app
                 }

                   fmt.Println(res)
    }
    }
// ******************** THE BELOW CODE IS THE ANALYTICS MODULE , AUTHOR: SAI RAM ********************

func Create(date string) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
	}

	db.Query("insert into analytics (date) values (?);", date)
	fmt.Println("saisai")
	//INSERT INTO `new_schema`.`output` (`date_out`) VALUES ('11');
	defer db.Close()
}
func Ep1stime(date string) (avg_time float64) {

	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())

	}
	defer db.Close()

	rows, err := db.Query("SELECT RESPTIME FROM EP1_RESP where DATE = ?;", date)
	if err != nil {

		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}

	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep1stime  = ? WHERE date = ?;", avg_time, date)

	return
}
func Ep1sttime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP1_RESP where (DATE = ? and STATUS = 'success') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep1sttime  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep1sftime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP1_RESP where (DATE = ? and STATUS = 'failure') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep1sftime  = ? WHERE date = ?;", avg_time, date)
	return
}

func Ep2stime(date string) (avg_time float64) {

	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())

		defer db.Close()
	}

	rows, err := db.Query("SELECT RESPTIME FROM EP2_RESP where DATE = ?;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep3_s_time float64
		if err := rows.Scan(&ep3_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep3_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep2stime  = ? WHERE date = ?;", avg_time, date)

	return
}

func Ep2sttime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP2_RESP where (DATE = ? and STATUS = 'sucess') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep2sttime  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep2sftime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP2_RESP where (DATE = ? and STATUS = 'fail') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep2sftime  = ? WHERE date = ?;", avg_time, date)
	return
}

func Ep3stime(date string) (avg_time float64) {

	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())

		defer db.Close()
	}

	rows, err := db.Query("SELECT RESPTIME FROM EP3_RESP where DATE = ?;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep3_s_time float64
		if err := rows.Scan(&ep3_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep3_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep3stime  = ? WHERE date = ?;", avg_time, date)

	return
}
func Ep3sttime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP3_RESP where (DATE = ? and TESTRESULT = 'pass') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep3sttime  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep3sftime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP3_RESP where (DATE = ? and TESTRESULT = 'fail'  and ACTID != '') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep3sftime  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep4stime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("SELECT RESPTIME FROM EP4_RESP where DATE = ?;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep4_s_time float64
		if err := rows.Scan(&ep4_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep4_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	fmt.Println(avg_time)
	db.Query("UPDATE analytics SET ep4stime  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep4sttime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP4_RESP where (DATE = ? and TESTRESULT = 'pass') ", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep4sttime  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep4sftime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP4_RESP where (DATE = ? and TESTRESULT = 'fail') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	_, err1 := db.Query("UPDATE analytics SET ep4sftime  = ? WHERE date = ?;", avg_time, date)
	if err1 != nil {
		log.Fatal(err)
	}

	return
}
func Ep5stime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("SELECT RESPTIME FROM EP5_RESP where DATE = ?;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep4_s_time float64
		if err := rows.Scan(&ep4_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep4_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	fmt.Println(avg_time)
	db.Query("UPDATE analytics SET ep5time  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep5sttime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP5_RESP where (DATE = ? and TESTRES != 'fail') ", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = (totel_time / incr)
	//fmt.Println(avg_time)
	_, err1 := db.Query("UPDATE analytics SET ep5sttime  = ? WHERE date = ?;", avg_time, date)
	if err1 != nil {
		log.Fatal(err1)
	}
	return
}
func Ep5sftime(date string) (avg_time float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())
		defer db.Close()
	}
	rows, err := db.Query("select RESPTIME FROM EP5_RESP where (DATE = ? and TESTRES = 'fail') ;", date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	totel_time := 0.0
	incr := 0.0
	for rows.Next() {
		var ep1_s_time float64
		if err := rows.Scan(&ep1_s_time); err != nil {
			log.Fatal(err)
		}
		totel_time = totel_time + ep1_s_time
		incr = incr + 1
	}
	avg_time = totel_time / incr
	db.Query("UPDATE analytics SET ep5sftime  = ? WHERE date = ?;", avg_time, date)
	return
}
func Ep3Acc(date_in string) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())

	}
	defer db.Close()

	cont, err := db.Query("SELECT COUNT(TESTRESULT) FROM EP3_RESP WHERE (date = ?);", date_in)
	if err != nil {
		log.Print(err.Error())
	}
	defer cont.Close()
	var tt float64
	for cont.Next() {
		var cont1 float64
		cont.Scan(&cont1)
		fmt.Println(cont1)
		tt = cont1
		print(tt)
	}
	con, err := db.Query("SELECT COUNT(TESTRESULT) FROM EP3_RESP WHERE (date = ? and TESTRESULT = 'pass');", date_in)
	if err != nil {
		log.Print(err.Error())
	}
	defer con.Close()
	var dt float64
	for con.Next() {
		var cont1 float64
		con.Scan(&cont1)
		fmt.Println(cont1)
		dt = cont1
		print(dt)
	}
	accuracy := dt / tt * 100
	fmt.Println("per", accuracy)
	_, err1 := db.Query("update analytics set ep3acctt= ? where date= ?", accuracy, date_in)
	if err1 != nil {
		log.Print(err.Error())

	}
}
func Ep4Acc(date_in string) (attt, aft float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())

	}
	defer db.Close()

	cont, err := db.Query("SELECT COUNT(EXPOP) FROM EP4_RESP WHERE (date = ? and EXPOP = '1' and TESTRESULT ='pass');", date_in)
	if err != nil {
		log.Print(err.Error())
	}
	defer cont.Close()
	var tt float64
	for cont.Next() {
		var cont1 float64
		cont.Scan(&cont1)
		fmt.Println(cont1)
		tt = cont1
		fmt.Print("tt", tt)
	}

	con, err := db.Query("SELECT COUNT(EXPOP) FROM EP4_RESP WHERE (date = ? and EXPOP = '1');", date_in)
	if err != nil {
		log.Print(err.Error())
	}
	defer con.Close()
	var dt float64
	for con.Next() {
		var cont1 float64
		con.Scan(&cont1)
		fmt.Println(cont1, "HEHEHE")
		dt = cont1
		print("SAdt", dt)
	}
	cont1, err := db.Query("SELECT COUNT(EXPOP) FROM EP4_RESP WHERE (date = ? and EXPOP = '0' and TESTRESULT ='pass');", date_in)
	if err != nil {
		log.Print(err.Error())

	}

	defer cont1.Close()
	var ff float64
	for cont1.Next() {
		var cont float64
		cont1.Scan(&cont)
		fmt.Println(cont)
		ff = cont
		print(ff, "ff")
	}

	con2, err := db.Query("SELECT COUNT(EXPOP) FROM EP4_RESP WHERE (date = ? and TESTRESULT = 'pass');", date_in)
	if err != nil {
		log.Print(err.Error())

	}

	defer con2.Close()
	var df float64
	for con2.Next() {
		var cont float64
		con2.Scan(&cont)
		fmt.Println(cont)
		df = cont
		if df == 0 {
			df = 1
		}
		if dt == 0 {
			dt = 1
		}
		print(df, "df")
	}
	attt = tt / dt * 100
	aft = ff / df * 100
	fmt.Println(attt, "percent")
	_, err1 := db.Query("update analytics set  ep4accft =?, ep4acctt = ? where date =?", aft, attt, date_in)
	if err1 != nil {
		log.Print(err.Error())

	}
	return
}
func Ep5Acc(date_in string) (attt, aft float64) {
	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
	if err != nil {
		log.Print(err.Error())

	}
	defer db.Close()

	cont, err := db.Query("SELECT COUNT(EXPRES) FROM EP5_RESP WHERE (date = ? and EXPRES = 'ok' and ACTRES ='ok');", date_in)
	if err != nil {
		log.Print(err.Error())
	}
	defer cont.Close()
	var tt float64
	for cont.Next() {
		var cont1 float64
		cont.Scan(&cont1)
		fmt.Println(cont1)
		tt = cont1
		fmt.Print("tt", tt)
	}

	con, err := db.Query("SELECT COUNT(EXPRES) FROM EP5_RESP WHERE (date = ? and EXPRES = 'ok');", date_in)
	if err != nil {
		log.Print(err.Error())
	}
	defer con.Close()
	var dt float64
	for con.Next() {
		var cont1 float64
		con.Scan(&cont1)
		fmt.Println(cont1, "HEHEHE")
		dt = cont1
		print("SAdt", dt)
	}
	cont1, err := db.Query("SELECT COUNT(EXPRES) FROM EP5_RESP WHERE (date = ? and EXPRES = 'failed' and ACTRES ='failed');", date_in)
	if err != nil {
		log.Print(err.Error())

	}

	defer cont1.Close()
	var ff float64
	for cont1.Next() {
		var cont float64
		cont1.Scan(&cont)
		fmt.Println(cont)
		ff = cont
		print(ff, "ff")
	}

	con2, err := db.Query("SELECT COUNT(EXPRES) FROM EP5_RESP WHERE (date = ? and EXPRES = 'failed');", date_in)
	if err != nil {
		log.Print(err.Error())

	}

	defer con2.Close()
	var df float64
	for con2.Next() {
		var cont float64
		con2.Scan(&cont)
		fmt.Println(cont)
		df = cont
		if df == 0 {
			df = 1
		}
		if dt == 0 {
			dt = 1
		}
		print(df, "df")
	}
	attt = tt / dt * 100
	aft = ff / df * 100
	fmt.Println(attt, "percent")
	_, err1 := db.Query("update analytics set  ep5accft =?, ep5acctt = ? where date =?", aft, attt, date_in)
	if err1 != nil {
		log.Print(err.Error())

	}
	return
}


func ep3MultiReq(w http.ResponseWriter,r *http.Request){

fmt.Fprintf(w,"starting of EP3 Multi req")
                 type Db_values struct
                 {
                   Uid string
                   image_path string

                 }
                  var Values Db_values


                 //connect to local database

                  db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat");

                 if err != nil {
                   log.Print(err.Error())
                 }

                 defer db.Close()

                 // pass a select query to reteieve uid and image path


                 results, err := db.Query("SELECT * FROM regtable;")
                 if err != nil {
                   panic(err.Error()) // proper error handling instead of panic in your app
                 }
                 for results.Next() {

                   // for each row, scan the result into our tag composite object
                   err = results.Scan(&Values.Uid,&Values.image_path)
                   if err != nil {
                     panic(err.Error()) // proper error handling instead of panic in your app
                   }
                               // and then print out the tag's Name attribute
                   //var render_data string = uid
               fmt.Println("Uid: ",Values.Uid,"image_path:",Values.image_path)

                 var ref_img_path1 string = Values.image_path
                   wg.Add(1)

                go ep3GoReq(ref_img_path1)
                wg.Wait()
            }


                             type multires struct{
                               nr int
                               avgrt float64
                             }


                             //send the analytics of Multiple Request

                             t := time.Now()
                             a:= strconv.Itoa(t.Year())
                             b:= strconv.Itoa(int(t.Month()))
                             c:= strconv.Itoa(t.Day())
                             d:= a+"-"+b+"-"+c


                            //fmt.Println(d.Format(time.RFC3339))
                             resul, err := db.Query("SELECT Count(*),Max(resptime) FROM MultiReqEP3 WHERE Date = ?;",d)
                             if err != nil {
                               panic(err.Error()) // proper error handling instead of panic in your app
                             }

                             var MmmRrrr multires
                             for resul.Next(){

                               err = resul.Scan(&MmmRrrr.nr,&MmmRrrr.avgrt)
                               fmt.Println()
                               if err != nil {
                                 panic(err.Error()) // proper error handling instead of panic in your app
                               }

                             }
                             resu, err := db.Query("INSERT INTO EP3M_RESP(DATE,NO_OF_RESP,TIME) values(?,?,?);",d,MmmRrrr.nr,MmmRrrr.avgrt)
                             if err != nil {
                               panic(err.Error()) // proper error handling instead of panic in your app
                             }

                             fmt.Println(resu)

                             fmt.Fprintf(w,"done with Ep3 multiple req")


          }
func ep3GoReq(ep3imgPath string){
         defer wg.Done()
         type Db_values struct
                 {
                   Uid string
                   image_path string

                 }
                 var Values Db_values
         var expop string
                  var actop string

                  var status string
                   var resp_code int

 extraParams := map[string]string{
                                    }

   url := "http://staging.vishwamcorp.com/v2/face_lookup_ios"

   request, err := newfileUploadRequest(url, extraParams, "image", ep3imgPath)

     request.Header.Add("Content-Type", "multipart/form-data")
      if err != nil {
          log.Fatal(err)
                                 }
                                 client := &http.Client{}
                                   startx1 := time.Now()
                                 resp, err := client.Do(request)


                                 tx1 := time.Now()
                                 elapsedx1 := tx1.Sub(startx1)
                                 var resptime_ep3 float64 = elapsedx1.Seconds()
                                 fmt.Println("The  Response Time for sending multiple requests to EP3:",elapsedx1)

                                 if err != nil {
                                    log.Fatal(err)
                                 } else {
                                   body_values_ep3, err := ioutil.ReadAll(resp.Body)
                                     //body := &bytes.Buffer{}
                                   //_, err := body.ReadFrom(resp.Body)
                                    //var m map[string]string

                                    //sep3, err := getep3([]byte(body_values_ep3))
                                    //mmm,err:= json.Marshal(body_values_ep3)
                                    if err != nil {
                                       log.Fatal(err)
                                    }
                                    var ep3_resp map[string]string
                                    json.Unmarshal(body_values_ep3, &ep3_resp)
                                    fmt.Println(ep3_resp["userId"])
                                    expop = Values.Uid
                                    actop = ep3_resp["userId"]

                                     var testresultep3 string
                                    //fmt.Println("VERY VERY IMPORTANT:",body_values_ep3)
                                    if expop == actop{
                                      testresultep3 = "pass"
                                    } else{
                                      testresultep3 = "fail"
                                    }
                                    resp.Body.Close()
                                    fmt.Println(resp.StatusCode)
                                    fmt.Println(resp.Header)
                                    fmt.Println(testresultep3)
                                    //fmt.Println(body_values_ep3)
                                 }


                                  if resp.StatusCode == 200{
                                    fmt.Println("UserId is successfully retrieved")
                                    status = "sucess"
                                    resp_code = resp.StatusCode

                                  }else{
                                    status = "fail"
                                    resp_code = resp.StatusCode
                                  }
                                  fmt.Println(status,resp_code)
                                  t:= time.Now()

                                   t.Format(time.RFC3339)
                                    db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
                                   res, err := db.Query("INSERT INTO MultiReqEP3(DATE,RESPONSE,RESPTIME,STATUS) values(?,?,?,?);",t,resp_code,resptime_ep3,status)
                                      if err != nil {
                                        panic(err.Error()) // proper error handling instead of panic in your app
                                      }

                                        fmt.Println(res)





}
    func analytics_home(w http.ResponseWriter,r *http.Request)  {


    	fmt.Fprintf(w,"Welcome to FR API Analytics Platform")


    }


    func runAnalytics(w http.ResponseWriter,r *http.Request){

    	date := "2018-08-01"

    	fmt.Fprintf(w,"Hurray  Analytics functions are running succesfully")
    	Create(date)
    	Ep1stime(date)
    	Ep1sttime(date)
    	Ep1sftime(date)
    	Ep2stime(date)
    	Ep3stime(date)
    	Ep4stime(date)
    	Ep2sttime(date)
    	Ep2sftime(date)
    	Ep3sttime(date)
    	Ep3sftime(date)
    	Ep4sttime(date)
    	Ep4sftime(date)
    	Ep4Acc(date)
    	Ep3Acc(date)
    	Ep5stime(date)
    	Ep5sttime(date)
    	Ep5sftime(date)
    	Ep5Acc(date)


    fmt.Fprintf(w,"I am done Analyzing results of FR API reponses. Thank you.....!!!")

    }

    func Epavg(w http.ResponseWriter,r *http.Request) {

    	db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
    	if err != nil {
    		log.Print(err.Error())

    	}
    	defer db.Close()
    	var a string = "ep1stime"
    	//fmt.Println("colum name")
    	//fmt.Scan(&a)

    	quer := "SELECT AVG(" + a + ") FROM analytics; "
    	cont1, err := db.Query(quer)
    	if err != nil {
    		log.Print(err.Error())

    	}
    	defer cont1.Close()
    	var ff float32
    	for cont1.Next() {
    		var cont float32
    		cont1.Scan(&cont)
    		//fmt.Println(cont)
    		ff = cont
    		fmt.Println(ff, "ff")
    	}

    	fmt.Fprintf(w,"Yeah, average function is exected successfully")

    }





    //****** Below are the functions related to Client to DB Module ********\\



    var tpl *template.Template


    func init() {
        tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
    }



    func index(w http.ResponseWriter, r *http.Request) {
        tpl.ExecuteTemplate(w, "index.gohtml", nil)



    }

    func processor(w http.ResponseWriter, r *http.Request) {
        //code for registration
        if r.Method != "POST" {
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }
        usid := r.FormValue("uid")
        file, _, err := r.FormFile("image")
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
            return
        }
        img, _, err := image.Decode(file)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
            return
        }
        m := resize.Resize(300, 300, img, resize.Lanczos3)

        var image_path string = "./images/"
        image_path += usid + ".jpg"

        out, err := os.Create(image_path)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        defer out.Close()
        //convertToPNG(out, file)

        _, err = io.Copy(out, file)
        if err != nil {
            log.Fatal(err)
        }
        file.Close()
        fmt.Println("Success")

        // Encode into jpeg http://blog.golang.org/go-image-package
        err = jpeg.Encode(out, m, nil)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }

        fmt.Print(out)
        fmt.Println(usid)




        aws_access_key_id := "AKIAJ7SZ44TTQUEUNCAA"
        aws_secret_access_key := "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"
        token := ""
        creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

        _, err1 := creds.Get()
        if err1 != nil {
            // handle error
        }

        cfg := aws.NewConfig().WithRegion("ap-south-1").WithCredentials(creds)

        svc := s3.New(session.New(), cfg)

        file2, err := os.Open(image_path)

        if err != nil {
            // handle error
        }

        defer file2.Close()

        fileInfo, _ := file2.Stat()

        size := fileInfo.Size()

        buffer := make([]byte, size) // read file content to buffer

        file2.Read(buffer)

        fileBytes := bytes.NewReader(buffer)
        fileType := http.DetectContentType(buffer)
        path := "/media/" + file2.Name()

        params := &s3.PutObjectInput{
            Bucket: aws.String("sukshi1"),
            Key: aws.String(path),
            Body: fileBytes,
            ContentLength: aws.Int64(size),
            ContentType: aws.String(fileType),
        }

        resp, err := svc.PutObject(params)
        if err != nil {
            // handle error
        }

        fmt.Printf("response %s", awsutil.StringValue(resp))

        var cloud_reg string = "https://s3.ap-south-1.amazonaws.com/sukshi1/media/images/"
        cloud_reg += usid + ".jpg"

        db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
        if err != nil {
            panic(err.Error())
        }
        defer db.Close()
        fmt.Println("succesfully connected to db")
        res, err := db.Query("INSERT INTO regtable(usid,image_path) values(?,?);",usid,cloud_reg)
        if err != nil {
            log.Println("User ID already registered")
            //http.Error(w, "User already registered", 500)
            //http.HandleFunc("/", index)
            tpl.ExecuteTemplate(w, "index.gohtml", nil)
            tpl.ExecuteTemplate(w, "responsepage.gohtml", nil)

        }
        if err == nil{
            tpl.ExecuteTemplate(w, "index.gohtml", nil)
            tpl.ExecuteTemplate(w, "processor.gohtml", nil)

        }

        fmt.Println(res)


        //executing respose template-
        //tpl.ExecuteTemplate(w, "processor.gohtml", nil)
    }


    func presentimg(w http.ResponseWriter, r *http.Request){
        if r.Method != "POST" {
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }
        usid2 := r.FormValue("uid2")

        //processing image 1
        file2, _, err := r.FormFile("image2")

        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
            return
        }
        img, _, err := image.Decode(file2)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
            return
        }
        m := resize.Resize(300, 300, img, resize.Lanczos3)

        var image_path2 string = "./presentimages/"
        image_path2 += usid2 + "1"+".jpg"

        out, err := os.Create(image_path2)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        defer out.Close()
        //convertToPNG(out, file2)

        _, err = io.Copy(out, file2)
        if err != nil {
            log.Fatal(err)
        }
        file2.Close()
        fmt.Println("Success")

        // Encode into jpeg http://blog.golang.org/go-image-package
        err = jpeg.Encode(out, m, nil)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }

        fmt.Print(out)

        //uploading image1 to s3 bucket-------------------------------------------------------------------------


        aws_access_key_id := "AKIAJ7SZ44TTQUEUNCAA"
        aws_secret_access_key := "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"
        token := ""
        creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

        _, err1 := creds.Get()
        if err1 != nil {
            // handle error
        }

        cfg := aws.NewConfig().WithRegion("ap-south-1").WithCredentials(creds)

        svc := s3.New(session.New(), cfg)

        file4, err := os.Open(image_path2)

        if err != nil {
            // handle error
        }

        defer file4.Close()

        fileInfo1, _ := file4.Stat()

        size := fileInfo1.Size()

        buffer := make([]byte, size) // read file content to buffer

        file4.Read(buffer)


        fileBytes := bytes.NewReader(buffer)
        fileType := http.DetectContentType(buffer)
        path := "/presimages/" + file4.Name()

        params := &s3.PutObjectInput{
            Bucket: aws.String("sukshi1"),
            Key: aws.String(path),
            Body: fileBytes,
            ContentLength: aws.Int64(size),
            ContentType: aws.String(fileType),
        }

        resp, err := svc.PutObject(params)
        if err != nil {
            // handle error
        }

        fmt.Printf("response %s", awsutil.StringValue(resp))


        var cloud_pres string = "https://s3.ap-south-1.amazonaws.com/sukshi1/presimages/presentimages/"

        cloud_pres += usid2+".jpg"


        db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
        if err != nil {
            panic(err.Error())
        }
        defer db.Close()
        fmt.Println("succesfully connected to db")
        res, err := db.Query("INSERT INTO presimgtable(username,pres_img_path) values(?,?);",usid2,cloud_pres)
        if err != nil {
            tpl.ExecuteTemplate(w, "index.gohtml", nil)
            tpl.ExecuteTemplate(w, "gestureform.gohtml", nil)
        }
        if err == nil {
            tpl.ExecuteTemplate(w, "index.gohtml", nil)
            tpl.ExecuteTemplate(w, "successresp.gohtml", nil)
        }

        fmt.Println(res)


        //tpl.ExecuteTemplate(w, "responsepage.gohtml", nil)

    }

    func gestureimg(w http.ResponseWriter, r *http.Request){
        if r.Method != "POST" {
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }
        //getting uid

        usid3 := r.FormValue("uid3")

        //processing gesture image

        file4, _, err := r.FormFile("imageges")

        if err != nil {
            log.Println(err)
            fmt.Println("error1")
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
            return
        }
        img, _, err := image.Decode(file4)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
            return
        }
        m := resize.Resize(300, 300, img, resize.Lanczos3)

        var image_path4 string = "./gestureimages/"
        image_path4 += usid3 + "_gesture"+".jpg"

        out, err := os.Create(image_path4)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        defer out.Close()
        //convertToPNG(out, file4)

        _, err = io.Copy(out, file4)
        if err != nil {
            log.Fatal(err)
        }
        file4.Close()
        fmt.Println("Success")

        // Encode into jpeg http://blog.golang.org/go-image-package
        err = jpeg.Encode(out, m, nil)
        if err != nil {
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }

        fmt.Print(out)

        // getting gesture int

        gesint := r.FormValue("gint")

        gint,_ := strconv.Atoi(gesint)

        fmt.Print(gint)

        //getting true or false
        truerfalse := r.FormValue("trf")

        trf, _ := strconv.Atoi(truerfalse)

        fmt.Print(trf)

        //uploading gesture image into cloud

        aws_access_key_id := "AKIAJ7SZ44TTQUEUNCAA"
        aws_secret_access_key := "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"
        token := ""
        creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

        _, err1 := creds.Get()
        if err1 != nil {
            // handle error
        }

        cfg := aws.NewConfig().WithRegion("ap-south-1").WithCredentials(creds)

        svc := s3.New(session.New(), cfg)

        file6, err := os.Open(image_path4)

        if err != nil {
            // handle error
        }

        defer file6.Close()

        fileInfo1, _ := file6.Stat()

        size := fileInfo1.Size()

        buffer := make([]byte, size) // read file content to buffer

        file6.Read(buffer)


        fileBytes := bytes.NewReader(buffer)
        fileType := http.DetectContentType(buffer)
        path := "/gesimages/" + file6.Name()

        params := &s3.PutObjectInput{
            Bucket: aws.String("sukshi1"),
            Key: aws.String(path),
            Body: fileBytes,
            ContentLength: aws.Int64(size),
            ContentType: aws.String(fileType),
        }

        resp, err := svc.PutObject(params)
        if err != nil {
            // handle error
        }

        fmt.Printf("response %s", awsutil.StringValue(resp))

        var cloud_ges string = "https://s3.ap-south-1.amazonaws.com/sukshi1/gesimages/gestureimages/"

        cloud_ges += usid3 + ".jpg"


        db, err := sql.Open("mysql", "sukshi16:sukshi16@tcp(fat.cre9qo68vgqh.ap-south-1.rds.amazonaws.com)/fat")
        if err != nil {
            panic(err.Error())
        }
        defer db.Close()
        fmt.Println("succesfully connected to db")
        res, err := db.Query("INSERT INTO gesimgtable(uname,gespath,gint,trf) values(?,?,?,?);",usid3,cloud_ges,gint,trf)
        if err != nil {
            tpl.ExecuteTemplate(w, "index.gohtml", nil)
            tpl.ExecuteTemplate(w, "gestureform.gohtml", nil)
        }
        if err ==nil {
            tpl.ExecuteTemplate(w, "presentform.gohtml", nil)
            tpl.ExecuteTemplate(w, "gesresponse.gohtml", nil)

        }

        fmt.Println(res)




        //tpl.ExecuteTemplate(w, "responsepage.gohtml", nil)
    }





    func main(){


        http.HandleFunc("/", index)
        http.HandleFunc("/process", processor)
        http.HandleFunc("/presentimg", presentimg)
        http.HandleFunc("/gestureimg", gestureimg)


     // ******************* below are handle func related to API Hitting Module ******************\\
        http.HandleFunc("/gesture_upload",ep5ReqResp)
        http.HandleFunc("/direct_match",ep4ReqResp)
        http.HandleFunc("/ref_upload",ep1ReqResp)
        http.HandleFunc("/retreive_ref",ep2ReqResp)
        http.HandleFunc("/face_lookup",ep3ReqResP)
        http.HandleFunc("/ref_multi",ep1MultipleReqResp)
        http.HandleFunc("/ep2_multi",ep2MultiReq)
        http.HandleFunc("/ep3_multi",ep3MultiReq)
        http.HandleFunc("/ep4_multi",ep4MultiReq)


        // ******************* below are handle func related to Analytics Module ******************\\
        http.HandleFunc("/analytics_home",analytics_home)
        http.HandleFunc("/analyze_results",runAnalytics)
        http.HandleFunc("/avg",Epavg)


        http.ListenAndServe(":8080",nil)






    }
