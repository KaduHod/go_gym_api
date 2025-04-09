package main

import "net/http"
type CustomServer struct {
    server *http.ServeMux
}
func (self CustomServer) handle(method string, path string,  handler func(w http.ResponseWriter, r *http.Request), middlewares ...Middleware) {
    var curr http.Handler
    if len(middlewares) != 0 {
        lastMiddleware := middlewares[len(middlewares) - 1]
        curr := lastMiddleware(http.HandlerFunc(handler))
        for i := len(middlewares) - 2; i >= 0; i-- {
            curr = middlewares[i](curr)
        }
    } else {
        curr = http.HandlerFunc(handler)
    }
    self.server.Handle(path, self.methodMiddleware(method)(curr))
}
func (self CustomServer) Get(path string, handler func(w http.ResponseWriter, r *http.Request), middlewares ...Middleware) {
    self.handle(http.MethodGet, path, handler, middlewares...)
}
func (self CustomServer) Post(path string, handler func(w http.ResponseWriter, r *http.Request), middlewares ...Middleware) {
    self.handle(http.MethodPost, path, handler, middlewares...)
}
func (self CustomServer) Put(path string, handler func(w http.ResponseWriter, r *http.Request), middlewares ...Middleware) {
    self.handle(http.MethodPut, path, handler, middlewares...)
}
func (self CustomServer) Delete(path string, handler func(w http.ResponseWriter, r *http.Request), middlewares ...Middleware) {
    self.handle(http.MethodDelete, path, handler, middlewares...)
}
func (self CustomServer) methodMiddleware(method string) Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Verifica se o método da requisição corresponde ao método esperado
            if r.Method != method {
                http.NotFound(w, r)
                return
            }
            // Se o método estiver correto, continua para o próximo handler
            next.ServeHTTP(w, r)
        })
    }
}
