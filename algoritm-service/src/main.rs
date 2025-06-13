#[macro_use]
extern crate rocket;
use rocket::serde::json::Json;

mod algoritm;

use algoritm::demo::{run_genetic_algorithm, BestSolutionResult};
use serde::Serialize;

#[derive(Serialize)]
struct ResultData {
    x: u32,
    fx: i32,
}

#[get("/")]
fn index() -> &'static str {
    "¡Hola, mundo desde Rocket!"
}

#[get("/run-algorithm")]
fn run_algorithm() -> Json<BestSolutionResult> {
    Json(run_genetic_algorithm())
}

#[launch]
fn rocket() -> _ {
    let config = rocket::Config {
        address: "0.0.0.0".parse().unwrap(),
        port: 8088,
        ..Default::default()
    };

    rocket::custom(config).mount("/", routes![index, run_algorithm])
}
