#[macro_use] extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "¡Hola, mundo desde Rocket! watch"
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![index])
        .configure(rocket::Config {
            address: std::net::Ipv4Addr::new(0, 0, 0, 0).into(),  // Escucha en 0.0.0.0
            port: 8088,  // Asegúrate de que coincida con el puerto expuesto
            ..rocket::Config::default()
        })
}
