use actix_web::{get, web, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn shipping_options_api() -> impl Responder {
    HttpResponse::Ok().body("Shipping Options API")
}

#[get("/health")]
async fn health() -> impl Responder {
    HttpResponse::Ok().body("OK")
}

#[get("/shipping-options")]
async fn list_shipping_options() -> impl Responder {
    HttpResponse::Ok().body("list shipping options endpoint")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(shipping_options_api)
            .service(health)
            .service(web::scope("/api/v1").service(list_shipping_options))
    })
    .bind(("127.0.0.1", 8081))?
    .run()
    .await
}
