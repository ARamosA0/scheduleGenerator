use crate::models::algoritm_config::AlgorithmConfig;
use rocket::serde::json::Json;

use crate::algoritm::demo::{run_genetic_algorithm, BestSolutionResult};
use serde::Serialize;

use crate::algoritm::schedule_main::{ejecutar_algoritmo_horario, execute_process};

#[get("/run-algorithm")]
pub async fn run_algorithm() -> Json<BestSolutionResult> {
    Json(run_genetic_algorithm())
}

#[get("/run-schedule")]
pub async fn run_schedule() {
    println!("\n--- INICIO EL ALGORITMO ---");
    ejecutar_algoritmo_horario()
}

#[post("/generar", format = "json", data = "<config>")]
pub async fn generar_horario(config: Json<AlgorithmConfig>) -> String {
    let mut config = config.into_inner();
    println!("Config recibida:\n{:#?}", config);

    execute_process(&config);
    "Algoritmo ejecutado con éxito".into()
}
