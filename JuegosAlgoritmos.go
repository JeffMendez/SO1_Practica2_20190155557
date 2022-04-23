import (    
    jugadoresNumeros[i] = random(1,7);
    "math/rand"
    "time"
)

func Ruleta (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    ganador := rand.Intn(jugadores + 1);
    if ganador == 0 { return 1; }
    else { return ganador; }
}

func Dados (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    jugadoresNumeros := make([]int,jugadores)
    for i,_ := range jugadores {
        jugadoresNumeros[i] = rand.Intn(1,7);
        if jugadoresNumeros[i] == 0 { jjugadoresNumeros[i] = jugadoresNumeros[i] + 1; }
    }
    mayor_actual := 0
    ganador := 0
    for i,_ := range jugadores { 
        if jugadoresNumeros[i] > mayor_actual {
            mayor_actual = jugadoresNumeros[i]
            ganador = i
        }
    }
    return ganador
}

func Dardos (jugadores int) int {
    tiroMayor := 0
    rand.Seed(time.Now().UnixNano())
    jugador := 0
    ganador := 0
    for i,_ := range jugadores {
        jugador = rand.Intn(jugadores + 1)
        if jugador == 0 { jugador += 1; }
        tiro := rand.Intn(6)
        if tiro == 0 { tiro += 1; }
        if tiro > tiroMayor { tiroMayor = tiro; ganador = jugador; }
        if tiro == 5 { return jugador; }
    }
    return ganador
}

func CartaMayor (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    for {
        carta := rand.Intn(52)
        ganador := rand.Intn(jugadores + 1)
        if ganador == 0 { ganador += 1; }
        if carta == 0 || carta == 13 || carta == 26 || carta == 39 { return ganador; }
    }
    return 1
}

func SillasMusicales (jugadores int) int {
    rand.Seed(time.Now().UnixNano())
    sillasNo := jugadores - 1
    sillas := make(int[],sillasNo)
    for {
        for i := 0; i < sillasNo {
            jugador := rand.Intn(jugadores + 1)
            if jugador == 0 { jugador = 1; }
            sillas[i] = jugador
        }
        if sillasNo == 1 { return sillas[0]; }
        else { sillasNo = sillasNo - 1; }
    }
    return 1
}
