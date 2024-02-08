module.exports = {
    HTML:function(user, ui, printer) {
        let head, menu, list_menu, list = ``;
        if (ui == 'main') {
            head = `복합기 관리 솔루션<br/>
            ${user} 관리자님 어서오세요.`;
            list_menu = `<div class="col-sm-4">
            <ul class="nav nav-pills flex-column">
                <li class="nav-item">
                    <a class="nav-link" href="/">관리 복합기</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link active" href="/service">서비스 내역</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="#">회원 관리</a>
                </li>
            </ul>
            <hr class="d-sm-none">
            </div>`;
            menu = `<li class="nav-item">
            <a class="nav-link" href="/logout_process">로그아웃</a>
            </li>`;
            list = `<div class="col-sm-8">
            <table class="table table-hover">
            <thead>
            <tr>
                <th scope="col">Center ID</th>
                <th scope="col">Printer ID</th>
                <th scope="col">Service ID</th>
                <th scope="col">-</th>
            </tr>
            </thead>
            <tbody>`;
            list = list + printer;
            list = list + `</tbody>
            </table>
            <button type="button" class="btn btn-primary" onClick="location.href='/service/enroll'">서비스 내역 등록</button>`
        } else if (ui == 'enroll') {
            list = `<div class="col-sm-10">
            <form action="http://localhost:3000/service/enroll_process" method="post">
            <table>
              <tr>
                <td><p>Printer ID</p></td>
                <td><p><input type="text" name="PID" placeholder="PID"></p></td>
              </tr>
              <tr>
                <td><p>Service ID</p></td>
                <td><p><input type="text" name="SID" placeholder="ID"></p></td>
              </tr>
              <tr>
                <td><p>소모품 수량</p></td>
                <td><p><input type="text" name="NUM" placeholder="NUM"></p></td>
              </tr>
              <tr>
                <td></td>
                <td><p>
                  <input type="submit" class="btn btn-primary" value="등록"/>
                  <button type="button" class="btn" onClick="location.href='/service'">취소</button>
                </p></td>
              </tr>
            </table>
            </form>
            </div>`;
            head = `${user}
            서비스 제공 내역 등록`;
            menu = `<li class="nav-item">
            <a class="nav-link" href="/logout_process">로그아웃</a>
            </li>`;
            list_menu = `<p></p>`; 
        }
      return UI(head, menu, list_menu, list);
    },list:function(id , queryObj){
      var list = ``;
      list = list + `<tr>`;
      list = list + `<td>${id}</td>`;
      list = list + `<td>${queryObj.pid}</td>`;
      list = list + `<td>${queryObj.serviceCode}</td>`;
      list = list + `<td>${queryObj.num}</td>`;
      list = list + `</tr>`;
      return list;
    }
  }
function UI(head, menu, list_menu, list) {
  return `
  <!DOCTYPE html>
  <html lang="en">
  <head>
    <title>Capstone Design</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js"></script>
  </head>
  <body>

  <div class="jumbotron text-center" style="margin-bottom:0">
    <h1>${head}</h1>
  </div>

  <nav class="navbar navbar-expand-sm bg-secondary navbar-dark">
    <a class="navbar-brand" href="/">복합기</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="collapsibleNavbar">
      <ul class="navbar-nav">
        ${menu}
      </ul>
    </div>
  </nav>

  <div class="container" style="margin-top:30px">
    <div class="row">
      ${list_menu}
      ${list}
      </div>
    </div>
  </div>
  </body>
  </html>
  `;
}