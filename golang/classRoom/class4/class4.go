package class4

import (
	"log"
	"net/http"
)

func Mainclass4() {
	/*content := `<html>
	    <header>
	        <title>This is my website !!</title>
	        <style>
	            th {
	                color: blue;
	                padding: 10px;
	            }
	            td {
	                text-align: center;
	                padding: 5px;
	            }
	        </style>
	    </header>
	    <body>
	        <img src='https://media.istockphoto.com/photos/european-short-haired-cat-picture-id1072769156?k=20&m=1072769156&s=612x612&w=0&h=k6eFXtE7bpEmR2ns5p3qe_KYh098CVLMz4iKm5OuO6Y=' height="300px"/>
	        <table border="1">
	            <thead>
	                <tr>
	                    <th>Emp. no.</th>
	                    <th>Emp. name</th>
	                </tr>
	            </thead>
	            <tbody>
	                <tr>
	                    <td>1007348</td>
	                    <td>Voramet</td>
	                </tr>
	            </tbody>
	        </table>
	    </body>
	</html>`*/
	fs := http.FileServer(http.Dir("./classRoom/class4/files"))
	/*handler := http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		w.Write([]byte(content))
	})*/
	http.Handle("/", fs)
	//err := http.ListenAndServe("0.0.0.0:3000", handler)
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		log.Println("cannot start server", err)
	}
}
