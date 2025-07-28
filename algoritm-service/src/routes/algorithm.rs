use crate::models::algoritm_config::RawData;
use crate::services::process_data::format_json;
use rocket::serde::json::Json;

use crate::algoritm::demo::{run_genetic_algorithm, BestSolutionResult};

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
pub async fn generar_horario(config: Json<RawData>) -> String {
    let config = config.into_inner();
    let formated_config = format_json(config);
    let result = execute_process(&formated_config);
    println!("RESULTADO ALGORITMO: \n{:#?}", result);
    "Algoritmo ejecutado con éxito".into()
}
