package main

import (
    pb "201901557/juegos/pb"
    "context"
    "math/rand"
    "time"
    "net"
    "google.golang.org/grpc"
    "fmt"
    "log"
    "github.com/optiopay/kafka/v2"
    "github.com/optiopay/kafka/v2/proto"
)

type servidorJuegos struct {
    pb.UnimplementedJuegosServer
}

func Ruleta (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    ganador := rand.Intn(jugadores + 1);
    if ganador == 0 { return 1; }
    return ganador
}

func Dados (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    jugadoresNumeros := make([]int,jugadores)
    for i,_ := range jugadoresNumeros {
        jugadoresNumeros[i] = rand.Intn(7);
        if jugadoresNumeros[i] == 0 { jugadoresNumeros[i] = 1; }
    }
    mayor_actual := 0
    ganador := 1
    for i,_ := range jugadoresNumeros {
        if jugadoresNumeros[i] > mayor_actual {
            mayor_actual = jugadoresNumeros[i]
            ganador = i + 1
        }
    }
    return ganador
}

func Dardos (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    for {
        tiro := rand.Intn(6)
        ganador := rand.Intn(jugadores + 1)
        if ganador == 0 { ganador = 1; }
        if tiro == 5 { return ganador; }
    }
    return 1
}

func CartaMayor (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    for {
        carta := rand.Intn(52)
        ganador := rand.Intn(jugadores + 1)
        if ganador == 0 { ganador = 1; }
        if carta == 0 || carta == 13 || carta == 26 || carta == 39 { return ganador; }
    }
    return 1
}

func SillasMusicales (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    sillasNo := jugadores - 1
    sillas := make([]int,sillasNo)
    for {
        for i := 0; i < sillasNo; i++ {
            jugador := rand.Intn(jugadores + 1)
            if jugador == 0 { jugador = 1; }
            sillas[i] = jugador
        }
        if sillasNo == 1 { 
            return sillas[0]; 
        } else { 
            sillasNo = sillasNo - 1; 
        }
    }
    return 1
}

func jugarJuego (juego, jugadores int) int {
    if juego == 1 { return Ruleta(jugadores);
    } else if juego == 2 { return Dados(jugadores);
    } else if juego == 3 { return Dardos(jugadores);
    } else if juego == 4 { return CartaMayor(jugadores);
    } else if juego == 5 { return SillasMusicales(jugadores); }
    return 1
}

func (s *servidorJuegos) Jugar(contexto context.Context, partida *pb.Partida) (*pb.Resultado, error) {
    juego := partida.GetJuego()
    if juego < 1 || juego > 5 { return &pb.Resultado{Error: 1}, nil; }
    ganador := jugarJuego(int(juego), int(partida.GetJugadores()))
    ganador = ganador + 0
    fmt.Println("-> Juego:", juego, " Ganador:", ganador)
    
    // Producir a servidor de kafka
    conf := kafka.NewBrokerConf("test-client")
    conf.AllowTopicCreation = true
    broker, err := kafka.Dial([]string{"my-cluster-kafka-bootstrap:9092"}, conf)
    if err != nil {
        fmt.Println("cannot connect to kafka cluster: %s", err)
    } else {
        defer broker.Close()
        producer := broker.Producer(kafka.NewProducerConf())
        msg := &proto.Message{Value: []byte(fmt.Sprintf("%d\t%d\t%d", juego, int(partida.GetJugadores()), ganador))}
        if _, err := producer.Produce("juegos", 0, msg); err != nil {
            fmt.Println("cannot produce message to %s:%d: %s", "juegos", 0, err)
        }
    }
    
    return &pb.Resultado{Error: 0}, nil
}

func main () {
    lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 5505))
    if err != nil { log.Fatalf("inicio fallido: %v",err); }
    fmt.Println("201901557 - Server gRPC iniciado")
    var opts []grpc.ServerOption
    grpcServer := grpc.NewServer(opts...)
    pb.RegisterJuegosServer(grpcServer, &servidorJuegos{})
    grpcServer.Serve(lis)
}
