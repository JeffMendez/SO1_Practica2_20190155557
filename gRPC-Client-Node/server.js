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

const env = require('node-getenv')

function aleatorio (min,max) {
    return Math.floor(Math.random() * (max - min) + min);
}

const servidor = Hapi.server({
    port: 4505,
    host: '0.0.0.0',
});

servidor.route({
    method: 'GET',
    path: '/',
    options: { cors: true },
    handler: function (solicitud, h) {
        return { "servidor": "funcionando.." };
    }
});

servidor.route({
    method: 'POST',
    path: '/jugar',
    options: { cors: true },
    handler: function (request, h) {
        // JUEGO
        var partida = request.payload;
        cliente = new juegos('0.0.0.0:5505', grpc.credentials.createInsecure());
        cliente.jugar(partida, function(error, resultado) {
            console.log("Peticion-http:", partida, "Resultado:", resultado, error)
        });
        return { "Juego": "Completado" };
    }
});

const init = async () => {
    console.log("Cliente iniciado")
    await servidor.start();
};

init();
