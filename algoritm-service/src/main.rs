#[macro_use]
extern crate rocket;
use rocket::serde::json::Json;

mod algoritm;
mod models;
mod routes;
mod services;

use algoritm::demo::{run_genetic_algorithm, BestSolutionResult};
use models::algoritm_config::AlgorithmConfig;
use routes::algorithm::{generar_horario, run_algorithm, run_schedule};
use serde::Serialize;

use crate::algoritm::schedule_main::{ejecutar_algoritmo_horario, execute_process};

#[derive(Serialize)]
struct ResultData {
    x: u32,
    fx: i32,
}

#[get("/")]
fn index() -> &'static str {
    "¡Hola, mundo desde Rocket!"
}

#[launch]
fn rocket() -> _ {
    let config = rocket::Config {
        address: "0.0.0.0".parse().unwrap(),
        port: 8088,
        ..Default::default()
    };

    rocket::custom(config).mount(
        "/",
        // routes![index, run_algorithm, run_schedule, generar_horario],
        routes![index, run_algorithm, run_schedule, generar_horario],
    )
}
