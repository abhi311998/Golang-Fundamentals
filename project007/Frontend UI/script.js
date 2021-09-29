var bottom5User_X, bottom5Bank_X, top5User_X, top5Bank_X, top20User_X, top20Bank_X;
var bottom5User_Y, bottom5Bank_Y, top5User_Y, top5Bank_Y, top20User_Y, top20Bank_Y;
var dummy;

const baseUrl = "http://192.168.43.244:10000/"

var serviceTabSelected = "Fetch";

function fetchDashboardData() {
    fetch(baseUrl+"dashboard", {method: 'GET'})
    .then((response) => { 
        response.json().then((data) => {
            // console.log("here: ", data);
            fillDashboardWithData(data)
        }).catch((err) => {
            console.log(err);
        })
    });
}

function getSum(data) {
    var sum = 0;
    for (let i = 0; i < data.length; i++) {
        sum += data[i].value;
    }
    return sum
}

function get_X_Y_values(data) {
    let x = [];
    let y = [];
    for(i=0;i<data.length;i++) {
        x.push(data[i].name)
        y.push(data[i].value)
    }
    // console.log("####################")
    // console.log(x, y)
    // console.log("####################")
    return [x, y]
}

function fillDashboardWithData(data) {
    // console.log("fillDashboardWithData:", data);
    
    document.getElementById("totalTxnAmnt").innerHTML = "$ "+ numberWithCommas(data["second"]["numRowsAffected"])
    document.getElementById("totalTxnAmnt2").innerHTML = "$ "+ numberWithCommas(getSum(data["first"][2]))
    document.getElementById("totalTxnAmnt3").innerHTML = "$ "+ numberWithCommas(getSum(data["first"][3]))

    dummy = get_X_Y_values(data["first"][0])
    top20User_X = dummy[0];
    top20User_Y = dummy[1];

    dummy = get_X_Y_values(data["first"][1])
    top20Bank_X = dummy[0];
    top20Bank_Y = dummy[1];

    dummy = get_X_Y_values(data["first"][2])
    top5User_X = dummy[0];
    top5User_Y = dummy[1];

    dummy = get_X_Y_values(data["first"][3])
    top5Bank_X = dummy[0];
    top5Bank_Y = dummy[1];

    dummy = get_X_Y_values(data["first"][4])
    bottom5User_X = dummy[0];
    bottom5User_Y = dummy[1];

    dummy = get_X_Y_values(data["first"][5])
    bottom5Bank_X = dummy[0];
    bottom5Bank_Y = dummy[1];

    prepareChart("myChart1", "bar", top20User_X, top20User_Y)
    prepareChart("myChart2", "doughnut", top5User_X, top5User_Y)
    prepareChart("myChart3", "doughnut", bottom5User_X, bottom5User_Y)

    // bank-chart-1
    prepareChart("bank-chart-1", "bar", top20Bank_X, top20Bank_Y)
    prepareChart("bank-chart-2", "doughnut", top5Bank_X, top5Bank_Y)
    prepareChart("bank-chart-3", "doughnut", bottom5Bank_X, bottom5Bank_Y)

}

function prepareChart(elementId, chartType, xValues, yValues) {
    var colorSet = ['#FF6633', '#FFB399', '#FF33FF', '#FFFF99', '#00B3E6', 
                    '#E6B333', '#3366E6', '#999966', '#99FF99', '#B34D4D',
                    '#80B300', '#809900', '#E6B3B3', '#6680B3', '#66991A', 
                    '#FF99E6', '#CCFF1A', '#FF1A66', '#E6331A', '#33FFCC'];

    var legendDisplay = false;
    if(chartType === "doughnut"){legendDisplay = true}
    
    var ctx = document.getElementById(elementId).getContext('2d'); // 2d context
    new Chart(ctx, {
        type: chartType,
        data: {
        labels: xValues,
        datasets: [{
            backgroundColor: colorSet,
            data: yValues
        }]
        },
        options: {
            title: {
                display: false,
                text: "Shares of Bank Transaction"
            },
            legend: {
                display: legendDisplay,
                position: 'right'
            }
        }
    });
}

function handleForm(e) {
    e.preventDefault()
    
    var sqlQueryResult = document.getElementById("sql-query-result");
    removeAllChildNodes(sqlQueryResult)

    if(serviceTabSelected === "Fetch") {
        removeAllChildNodes(sqlQueryResult);
        var formData = new FormData(document.getElementById("form-fetch"));
        fetchData(formData.get("option-fetch"), formData.get("textInput-fetch"));
    }else if(serviceTabSelected === "Insert") {
        removeAllChildNodes(sqlQueryResult);
        var formData = new FormData(document.getElementById("form-insert"));
        var strData = formData.get("textInput-insert");
        if(strData != ""){
            insertData(JSON.parse(strData));
        }
    }else if(serviceTabSelected === "Update") {
        removeAllChildNodes(sqlQueryResult);
        var formData = new FormData(document.getElementById("form-update"));
    }else if(serviceTabSelected === "Delete") {
        removeAllChildNodes(sqlQueryResult);
        var formData = new FormData(document.getElementById("form-update"));
        var strData = formData.get("textInput-update");
        if(strData != ""){
            updateData(JSON.parse(strData));
        }
    }
}

function fillTableData(data) {

    document.getElementById("services-fetch-tab").hidden = true;
    document.getElementById("services-insert-tab").hidden = true;
    document.getElementById("services-update-tab").hidden = true;
    document.getElementById("services-delete-tab").hidden = true;

    var sqlQueryResult = document.getElementById("sql-query-result");
    var table = document.createElement("table");
    for (var i = -1; i < data.length; i++) {
      var tr = document.createElement("tr");
      if (i == -1) {
        var timeStamp = document.createElement("th");
        var txnId = document.createElement("th");
        var upiId = document.createElement("th");
        var userName = document.createElement("th");
        var bankName = document.createElement("th");
        var txnAmount = document.createElement("th");

        timeStamp.appendChild(document.createTextNode("TimeStamp"));
        txnId.appendChild(document.createTextNode("Transaction Id"));
        upiId.appendChild(document.createTextNode("UPI Id"));
        userName.appendChild(document.createTextNode("User Name"));
        bankName.appendChild(document.createTextNode('Bank Name'));
        txnAmount.appendChild(document.createTextNode("Transaction Amount"));

        tr.appendChild(timeStamp);
        tr.appendChild(txnId);
        tr.appendChild(upiId);
        tr.appendChild(userName);
        tr.appendChild(bankName);
        tr.appendChild(txnAmount);
      } else {
        var timeStamp = document.createElement("td");
        var txnId = document.createElement("td");
        var upiId = document.createElement("td");
        var userName = document.createElement("td");
        var bankName = document.createElement("td");
        var txnAmount = document.createElement("td");

        timeStamp.appendChild(document.createTextNode(data[i]["timeStamp"]));
        txnId.appendChild(document.createTextNode(data[i]["txnId"]));
        upiId.appendChild(document.createTextNode(data[i]["upiId"]));
        userName.appendChild(document.createTextNode(data[i]["userName"]));
        bankName.appendChild(document.createTextNode(data[i]["bankName"]));
        txnAmount.appendChild(document.createTextNode(data[i]["txnAmount"]));

        tr.appendChild(timeStamp);
        tr.appendChild(txnId);
        tr.appendChild(upiId);
        tr.appendChild(userName);
        tr.appendChild(bankName);
        tr.appendChild(txnAmount);
      }
      table.appendChild(tr);
    }
    sqlQueryResult.append(table);
}

function fillBlock(data) {
    var sqlQueryResult = document.getElementById("sql-query-result");
    var h5 = document.createElement("h5");
    var text = document.createTextNode("Number of rows affected: " + data['numRowsAffected']);
    h5.appendChild(text);
    sqlQueryResult.append(h5)
}

function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}

function parametersBasedOnRequestType(that) {
    if (that.value === "data") {
        document.getElementById("secondRow-data").hidden = false;
        document.getElementById("secondRow-insert").hidden = true;
        document.getElementById("secondRow-update").hidden = true;
        document.getElementById("secondRow-delete").hidden = true;
    } else if(that.value === "insert") {
        document.getElementById("secondRow-data").hidden = true;
        document.getElementById("secondRow-insert").hidden = false;
        document.getElementById("secondRow-update").hidden = true;
        document.getElementById("secondRow-delete").hidden = true;
    } else if(that.value === "update") {
        document.getElementById("secondRow-data").hidden = true;
        document.getElementById("secondRow-insert").hidden = true;
        document.getElementById("secondRow-update").hidden = false;
        document.getElementById("secondRow-delete").hidden = true;
    } else if(that.value === "delete") {
        document.getElementById("secondRow-data").hidden = true;
        document.getElementById("secondRow-insert").hidden = true;
        document.getElementById("secondRow-update").hidden = true;
        document.getElementById("secondRow-delete").hidden = false;
    } else {
        document.getElementById("secondRow-data").hidden = true;
        document.getElementById("secondRow-insert").hidden = true;
        document.getElementById("secondRow-update").hidden = true;
        document.getElementById("secondRow-delete").hidden = true;
    }
}

async function fetchData(typeOfGetData, param) {
    let endpoint = "get";
    if(typeOfGetData === "username"){
        endpoint = "getUser/"+ param;
    }else if(typeOfGetData === "bankname"){
        endpoint = "getBank/"+ param;
    }
    
    fetch(baseUrl+endpoint, {method: 'GET'})
    .then((response) => { 
        response.json().then((data) => {
            console.log("here: ", JSON.stringify(data));
            fillTableData(data)
        }).catch((err) => {
            console.log(err);
        })
    });
}

async function insertData(dataToAdd) {
    let endpoint = "addTxn";
    console.log(dataToAdd, typeof(JSON.stringify(dataToAdd)))
    fetch(baseUrl+endpoint, {
        method: 'POST',
        headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataToAdd)
    })
    .then((response) => { 
        response.json().then((data) => {
            fillBlock(data)
            console.log("insert response returned: ", data)
        }).catch((err) => {
            console.log(err);
        })
    });
}

async function updateData(dataToAdd) {
    let endpoint = "addTxn";
    console.log(dataToAdd, typeof(JSON.stringify(dataToAdd)))
    fetch(baseUrl+endpoint, {
        method: 'POST',
        headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataToAdd)
    })
    .then((response) => { 
        response.json().then((data) => {
            fillBlock(data)
            console.log("insert response returned: ", data)
        }).catch((err) => {
            console.log(err);
        })
    });
}

async function deleteData(typeOfGetData, param) {
    let endpoint = "";
    if(typeOfGetData === "username"){
        endpoint = "delete/"+ param;
    }else if(typeOfGetData === "bankname"){
        endpoint = "delete/"+ param;
    }
    
    fetch(baseUrl+endpoint, {method: 'delete'})
    .then((response) => { 
        response.json().then((data) => {
            // console.log("here: ", data);
            fillBlock(data)
            console.log(data)
        }).catch((err) => {
            console.log(err);
        })
    });
}

function handleHomeNav() {

    function getEventTarget(e) {
        e = e || window.event;
        return e.target || e.srcElement; 
    }

    // alert("called", that.value)
    var ul = document.getElementById('test');
    ul.onclick = function(event) {
        var target = getEventTarget(event);
        var tab = target.innerHTML;
        if(tab === 'User'){
            document.getElementById("home-user-tab").hidden = false;
            document.getElementById("home-bank-tab").hidden = true;
            document.getElementById("home-dashboard-tab").hidden = true;
        } else if(tab === 'Bank'){
            document.getElementById("home-user-tab").hidden = true;
            document.getElementById("home-bank-tab").hidden = false;
            document.getElementById("home-dashboard-tab").hidden = true;
        } else {
            document.getElementById("home-user-tab").hidden = true;
            document.getElementById("home-bank-tab").hidden = true;
            document.getElementById("home-dashboard-tab").hidden = false;
        }
    };
}

function handleServicesNav(e) {
    e.preventDefault();
    function getEventTarget(e) {
        e = e || window.event;
        return e.target || e.srcElement; 
    }

    var ul = document.getElementById('test2');
    ul.onclick = function(event) {
        var sqlQueryResult = document.getElementById("sql-query-result");
        removeAllChildNodes(sqlQueryResult);
        
        var target = getEventTarget(event);
        var tab = target.innerHTML;
        serviceTabSelected = tab;
        if(tab === 'Insert'){
            document.getElementById("services-fetch-tab").hidden = true;
            document.getElementById("services-insert-tab").hidden = false;
            document.getElementById("services-update-tab").hidden = true;
            document.getElementById("services-delete-tab").hidden = true;
        } else if(tab === 'Update'){
            document.getElementById("services-fetch-tab").hidden = true;
            document.getElementById("services-insert-tab").hidden = true;
            document.getElementById("services-update-tab").hidden = false;
            document.getElementById("services-delete-tab").hidden = true;
        }else if(tab === 'Delete'){
            document.getElementById("services-fetch-tab").hidden = true;
            document.getElementById("services-insert-tab").hidden = true;
            document.getElementById("services-update-tab").hidden = true;
            document.getElementById("services-delete-tab").hidden = false;
        }else {
            document.getElementById("services-fetch-tab").hidden = false;
            document.getElementById("services-insert-tab").hidden = true;
            document.getElementById("services-update-tab").hidden = true;
            document.getElementById("services-delete-tab").hidden = true;
        }
    };
}

function numberWithCommas(x) {
    return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}