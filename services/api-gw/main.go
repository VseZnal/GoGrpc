package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	getService "Grpc/services/api-gw/protogw/getmem/proto"

	_ "github.com/subosito/gotenv"
	"log"
	"net/http"
)

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
	log.Printf("server listening at 8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
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
