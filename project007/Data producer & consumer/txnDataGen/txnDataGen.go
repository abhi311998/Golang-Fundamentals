package txnDataGen

import (
	"math/rand"
	"strings"
	"time"
	"strconv"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	// rand.Seed(time.Now().UnixNano())
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func getRandomDate() time.Time {
	max := time.Now().Unix()
	min := time.Now().AddDate(-1, 0, 0).Unix()

    delta := max - min

    sec := rand.Int63n(delta) + min
    return time.Unix(sec, 0)
}

func getBankName() string{
	// rand.Seed(time.Now().UnixNano())
	banks := []string{"Axis Bank", "Bandhan Bank", "CSB Bank", "Union Bank", "DCB Bank", 
						"Dhanlaxmi Bank", "Federal Bank", "HDFC Bank", "ICICI Bank", "IDBI Bank", 
						"IDFC Bank", "IndusInd Bank", "J&K Bank", "Karnataka Bank", "Karur Bank", 
						"Kotak Bank", "Nainital Bank", "RBL Bank", "Indian Bank", 
						"Tamilnad Bank", "YES Bank"}		
	
	return banks[rand.Intn(len(banks))]
}

func getUserName() string {
	// rand.Seed(time.Now().UnixNano())
	return "user" + strconv.Itoa(rand.Intn(50))
}

func getAmount() int64 {
	min, max := 100, 100000
	return int64(rand.Intn(max - min) + min)
}

func GenTxnData(n time.Duration) Transaction{
		rand.Seed(time.Now().UnixNano())
		usrname := getUserName()
		bnkname := getBankName()
		rndDate := getRandomDate().Format(timeFormat)
		txn := Transaction {
			TimeStamp: rndDate,
			UserName: usrname,
			BankName: bnkname,
			UpiId: usrname + "@" + strings.Split(bnkname, " ")[0],
			TxnId: randSeq(20),
			TxnAmount: getAmount(),
		}
		time.Sleep(time.Millisecond * n)
		// str := fmt.Sprintf("%v, %v, %v, %v, %v, %v", string(txn.TimeStamp), txn.TxnId, txn.UpiId, txn.UserName, txn.BankName, txn.TxnAmount)
		// log.Println("data: ", str)
		return txn
}