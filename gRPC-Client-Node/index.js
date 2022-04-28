const Hapi = require('@hapi/hapi');

var PROTO_PATH = __dirname + '/juegos.proto';
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
var juegos = protoDescriptor.Juegos;

const servidor = Hapi.server({
    port: 4505,
    host: '0.0.0.0',
});

servidor.route({
    method: 'GET',
    path: '/',
    options: { cors: true },
    handler: function (solicitud, h) {
        return { "cliente-gRPC": "201901557-JM" };
    }
});

servidor.route({
    method: 'POST',
    path: '/juego',
    options: { cors: true },
    handler: function (request, h) {
        var partida = { Juego: request.payload.game_id, Jugadores: request.payload.players } 
        cliente = new juegos('0.0.0.0:5505', grpc.credentials.createInsecure());
        cliente.jugar(partida, function(error, resultado) {
            console.log(" -> Respuesta del servidor: ", partida, resultado)
        });
        return { "Juego": "Procesado en Kafka", "Error": 0 };
    }
});

const init = async () => {
    console.log("201901557 - Cliente gRPC iniciado")
    await servidor.start();
};

init();
