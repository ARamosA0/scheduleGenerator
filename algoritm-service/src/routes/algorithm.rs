use crate::models::algoritm_config::RawData;
use crate::models::schedule_model::ScheduleResponse;
use crate::services::process_data::format_json;
use rocket::serde::json::Json;
use std::time::Instant;

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
pub async fn generar_horario(config: Json<RawData>) -> Json<ScheduleResponse> {
    let config = config.into_inner();
    let assignment_id = config.id.unwrap_or(0) as usize;
    let formated_config = format_json(config);
    let start = Instant::now();
    // println!("----------FORMATED CONFIG----------: \n{:#?}", formated_config);
    let mut result: ScheduleResponse = execute_process(&formated_config);
    let duration = start.elapsed();
    let seconds = duration.as_secs() as i64;
    result.time = seconds;
    result.assigment_id = assignment_id;
    // println!("RESULTADO ALGORITMO: \n{:#?}", result);
    Json(result)
}
