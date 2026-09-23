# Servicio de rutas óptimas

Servicio en Go que determina qué base de grúas puede atender un accidente con la menor distancia posible y devuelve el camino mínimo utilizando Dijkstra.

## Arquitectura

El algoritmo no depende del transporte HTTP:

```text
HTTP Handler
    |
Application Service
    |
Dijkstra
    |
Graph Domain
```

Esta separación permite reemplazar el origen del grafo o reutilizar el algoritmo sin modificar la lógica central.

## Endpoint

- `POST /api/v1/routes/optimal`: endpoint versionado.
- `POST /routes/optimal`: alias simple para ejecución local.
- `GET /health`: disponibilidad del servicio.

## Ejecución

```bash
make run
```

## Calidad

```bash
make test
make quality
make build
```

La solución contempla múltiples bases y devuelve un error controlado cuando el destino no puede alcanzarse desde ninguna de ellas.
