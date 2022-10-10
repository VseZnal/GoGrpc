package main

import (
	deleteService "Grpc/services/api-gw/protogw/deletemem/proto"
	getService "Grpc/services/api-gw/protogw/getmem/proto"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"regexp"

	createService "Grpc/services/api-gw/protogw/createmem/proto"
	_ "github.com/subosito/gotenv"
)

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := getService.RegisterGetmemServiceHandlerFromEndpoint(ctx, mux, "localhost:8083", opts)
	if err != nil {
		panic(err)
	}

	opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = deleteService.RegisterDeletememServiceHandlerFromEndpoint(ctx, mux, "localhost:8082", opts)
	if err != nil {
		panic(err)
	}

	opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = createService.RegisterCreatememServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", opts)
	if err != nil {
		panic(err)
	}

	gwServer := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8000")
	log.Fatalln(gwServer.ListenAndServe())

}

/*

var (
	// gRPC services
	createServerAdress = fmt.Sprintf("%s:%s", os.Getenv("CREATEMEM_HOST"), os.Getenv("CREATEMEM_PORT"))
	deleteServerAdress = fmt.Sprintf("%s:%s", os.Getenv("DELETEMEM_HOST"), os.Getenv("DELETEMEM_PORT"))
	getServerAdress    = fmt.Sprintf("%s:%s", os.Getenv("GETMEM_HOST"), os.Getenv("GETMEM_PORT"))
)

func main() {
	proxyAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	HTTPProxy(proxyAddr)
}

func HTTPProxy(proxyAddr string) {

	grpcGwMux := runtime.NewServeMux()

	//----------------------------------------------------------------
	// настройка подключений со стороны gRPC
	//----------------------------------------------------------------

	//Подключение к сервису createmem
	grpcCreateConn, err := grpc.Dial(
		createServerAdress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Filed to connect to Create service", err)
	}
	defer grpcCreateConn.Close()

	err = createService.RegisterCreatememServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcCreateConn,
	)
	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}

	//----------------------------------------------------------------
	//Подключение к сервису deletemem
	grpcDeleteConn, err := grpc.Dial(
		deleteServerAdress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Filed to connect to Delete service", err)
	}
	defer grpcDeleteConn.Close()

	err = deleteService.RegisterDeletememServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcDeleteConn,
	)
	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}

	//----------------------------------------------------------------
	//Подключение к сервису getmem
	grpcGetConn, err := grpc.Dial(
		getServerAdress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Filed to connect to Get service", err)
	}
	defer grpcGetConn.Close()

	err = getService.RegisterGetmemServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcDeleteConn,
	)
	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}

	//----------------------------------------------------------------
	//	TНастройка маршрутов с стороны RES
	//----------------------------------------------------------------
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", grpcGwMux)
	mux.HandleFunc("/", helloworld)

	fmt.Println("starting HTTP server at " + proxyAddr)
	log.Fatal(http.ListenAndServe(proxyAddr, mux))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL:", r.URL.String())
}

*/
