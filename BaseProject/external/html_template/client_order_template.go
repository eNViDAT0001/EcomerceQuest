package html_template

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	"strconv"
)

func GetClientOrderTemplate(order entities.Order) string {
	return `<!doctype html>
<html lang="en" xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office">

<head>
  <title> Notificatin [CES] </title>
  <!--[if !mso]><!-- -->
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <!--<![endif]-->
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <style type="text/css">
    #outlook a {
      padding: 0;
    }

    body {
      margin: 0;
      padding: 0;
      -webkit-text-size-adjust: 100%;
      -ms-text-size-adjust: 100%;
    }

    table,
    td {
      border-collapse: collapse;
      mso-table-lspace: 0pt;
      mso-table-rspace: 0pt;
    }

    img {
      border: 0;
      height: auto;
      line-height: 100%;
      outline: none;
      text-decoration: none;
      -ms-interpolation-mode: bicubic;
    }

    p {
      display: block;
      margin: 13px 0;
    }
  </style>
  <link href="https://fonts.googleapis.com/css?family=Roboto:400,500,700" rel="stylesheet" type="text/css">
  <style type="text/css">
    @import url(https://fonts.googleapis.com/css?family=Roboto:400,500,700);
  </style>
  <style type="text/css">
    @media only screen and (min-width:480px) {
      .mj-column-per-100 {
        width: 100% !important;
        max-width: 100%;
      }

      .mj-column-per-50 {
        width: 50% !important;
        max-width: 50%;
      }
    }
  </style>
  <style type="text/css">
    @media only screen and (max-width:480px) {
      table.mj-full-width-mobile {
        width: 100% !important;
      }

      td.mj-full-width-mobile {
        width: auto !important;
      }
    }
  </style>
  <style type="text/css">
    a,
    span,
    td,
    th {
      -webkit-font-smoothing: antialiased !important;
      -moz-osx-font-smoothing: grayscale !important;
    }
  </style>
</head>

<body style="background-color:#f3f3f5;">
<div style="display:none;font-size:1px;color:#ffffff;line-height:1px;max-height:0px;max-width:0px;opacity:0;overflow:hidden;"> Preview - Notification from CES </div>
<div style="background-color:#f3f3f5;">

  <div style="margin:0px auto;max-width:600px;">
    <table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="width:100%;">
      <tbody>
      <tr>
        <td style="direction:ltr;font-size:0px;padding:20px 0;text-align:center;">

          <div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;">
            <table border="0" cellpadding="0" cellspacing="0" role="presentation" style="vertical-align:top;" width="100%">
              <tr>
                <td style="font-size:0px;word-break:break-word;">

                  <div style="height:20px;"> &nbsp; </div>

                </td>
              </tr>
            </table>
          </div>

        </td>
      </tr>
      </tbody>
    </table>
  </div>

  <div style="background:#54595f;background-color:#54595f;margin:0px auto;border-radius:4px 4px 0 0;max-width:600px;">
    <table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="background:#54595f;background-color:#54595f;width:100%;border-radius:4px 4px 0 0;">
      <tbody>
      <tr>
        <td style="direction:ltr;font-size:0px;padding:20px 0;text-align:center;">

          <div style="margin:0px auto;max-width:600px;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="width:100%;">
              <tbody>
              <tr>
                <td style="direction:ltr;font-size:0px;padding:0px;text-align:center;">

                  <div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;">
                    <table border="0" cellpadding="0" cellspacing="0" role="presentation" style="vertical-align:top;" width="100%">
                      <tr>
                        <td style="font-size:0px;word-break:break-word;">

                          <div style="height:20px;"> &nbsp; </div>

                        </td>
                      </tr>
                      <tr>
                        <td align="center" style="font-size:0px;padding:10px 25px;padding-bottom:0px;word-break:break-word;">
                          <div style="font-family:Roboto, Helvetica, Arial, sans-serif;font-size:24px;font-weight:400;line-height:30px;text-align:center;color:#ffffff;">
                            <h1 style="margin: 0; font-size: 24px; line-height: normal; font-weight: 400;">Thank you for your order</h1>
                          </div>
                        </td>
                      </tr>
                      <tr>
                        <td align="center" style="font-size:0px;padding:10px 25px;padding-top:0;word-break:break-word;">
                          <div style="font-family:Roboto, Helvetica, Arial, sans-serif;font-size:14px;font-weight:400;line-height:20px;text-align:center;color:#aaaaaa;">
                            <p style="margin: 0;">Order Number: ` + strconv.Itoa(int(order.ID)) + ` | Order Date: ` + fmt.Sprintf("%d/%d/%d", order.CreatedAt.Day(), order.CreatedAt.Month(), order.CreatedAt.Year()) + ` </p>
                          </div>
                        </td>
                      </tr>
                      <tr>
                        <td style="font-size:0px;padding:10px 25px;word-break:break-word;">
                          <p style="border-top: solid 1px #999999; font-size: 1px; margin: 0px auto; width: 100%;">
                          </p>

                        </td>
                      </tr>
                      <tr>
                        <td align="left" style="font-size:0px;padding:10px 25px;word-break:break-word;">
                          <div style="font-family:Roboto, Helvetica, Arial, sans-serif;font-size:14px;font-weight:400;line-height:20px;text-align:left;color:white;">
                            <p style="margin: 0;">Hi ` + order.Name + `,<br> Thank you for your purchase. Please find your order summary below. </p>
                          </div>
                        </td>
                      </tr>
                    </table>
                  </div>

                </td>
              </tr>
              </tbody>
            </table>
          </div>

              <div style="margin:0px auto;max-width:600px;">
                <table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="width:100%;">
                  <tbody>
                    <tr>
                      <td style="direction:ltr;font-size:0px;padding:20px 0;text-align:center;">

                        <div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;">
                          <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%">
                            <tbody>
                              <tr>
                                <td style="vertical-align:top;padding:0px 25px;">
                                  <table border="0" cellpadding="0" cellspacing="0" role="presentation" style="background-color:#34393E;" width="100%">
                                    <tr>
                                      <td align="left" class="receipt-table" style="font-size:0px;padding:30px;word-break:break-word;">
                                        <table cellpadding="0" cellspacing="0" width="100%" border="0" style="color:#000000;font-family:Roboto, Helvetica, Arial, sans-serif;font-size:13px;line-height:22px;table-layout:auto;width:100%;border:none;">
                                          <tr>
                                            <th colspan="3" style="font-size: 20px; line-height: 30px; font-weight: 500; color: #fff; padding: 0px 0px 10px 0px; text-align: center; border-bottom: 1px solid #555;" align="center">Order summary </th>
                                          </tr>

                                          <tr>
                                            <td style="color: #ccc; font-size: 15px; line-height: 22px; font-weight: 400; word-break: normal; padding: 10px 0; text-align: left;" colspan="2" align="left">Total</td>
                                            <td style="color: #ccc; font-size: 15px; line-height: 22px; font-weight: 400; word-break: normal; padding: 10px 0; text-align: right;" align="right">` + fmt.Sprint(float64(order.Total)/((100-float64(order.Discount))/100.0)) + `$</td>
                                          </tr>
                                          <tr>
                                            <td style="color: #ccc; font-size: 15px; line-height: 22px; font-weight: 400; word-break: normal; padding: 10px 0; text-align: left;" colspan="2" align="left">Discount</td>
                                            <td style="color: #ccc; font-size: 15px; line-height: 22px; font-weight: 400; word-break: normal; padding: 10px 0; text-align: right;" align="right">` + strconv.Itoa(order.Discount) + `%</td>
                                          </tr>
                                          <tr>
                                            <td style="word-break: normal; color: #fff; font-size: 20px; line-height: 30px; border-top: 1px solid #555; font-weight: 500; padding: 10px 0px 0px 0px; text-align: left;" colspan="2" align="left">Summary</td>
                                            <td style="word-break: normal; color: #fff; font-size: 20px; line-height: 30px; border-top: 1px solid #555; font-weight: 500; text-align: right; padding: 10px 0px 0px 0px;" align="right">` + strconv.Itoa(order.Total) + `$</td>
                                          </tr>
                                        </table>
                                      </td>
                                    </tr>
                                  </table>
                                </td>
                              </tr>
                            </tbody>
                          </table>
                        </div>

                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

          <div style="margin:0px auto;max-width:600px;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="width:100%;">
              <tbody>
              <tr>
                <td style="direction:ltr;font-size:0px;padding:0px;text-align:center;">

                  <div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0;line-height:0;text-align:left;display:inline-block;width:100%;direction:ltr;">

                    <div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;">
                      <table border="0" cellpadding="0" cellspacing="0" role="presentation" style="vertical-align:top;" width="100%">
                        <tr>
                          <td align="left" style="font-size:0px;padding:10px 25px;word-break:break-word;">
                            <div style="font-family:Roboto, Helvetica, Arial, sans-serif;font-size:14px;font-weight:400;line-height:20px;text-align:left;color:#ffffff;">
                              <p style="margin: 0;">You're receiving this email because you made a purchase at CES. If you have any questions, contact us at <a href="#" style="color: #009BF9; text-decoration: none; word-break: normal;">dummyshop@gmail.com</a></p>
                            </div>
                          </td>
                        </tr>
                      </table>
                    </div>

                  </div>

                </td>
              </tr>
              </tbody>
            </table>
          </div>

        </td>
      </tr>
      </tbody>
    </table>
  </div>



</div>
</body>

</html>`
}
