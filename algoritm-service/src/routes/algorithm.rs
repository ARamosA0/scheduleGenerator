use crate::models::algoritm_config::RawData;
use crate::models::schedule_model::ScheduleResponse;
use crate::services::process_data::format_json;
use rocket::serde::json::Json;
use std::time::Instant;
use rocket::response::stream::{Event, EventStream};
use rocket::tokio::sync::mpsc;
use rocket::tokio::task;
use rocket::tokio::time::{sleep, Duration};

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
pub async fn generar_horario(config: Json<RawData>) -> EventStream![] {
    let config = config.into_inner();
    let assignment_id = config.id.unwrap_or(0) as usize;
    let formated_config = format_json(config);
    let start = Instant::now();
    // println!("----------FORMATED CONFIG----------: \n{:#?}", formated_config);
    // let mut result: ScheduleResponse = execute_process(&formated_config);
    let (tx, mut rx) = mpsc::channel::<String>(100);

    task::spawn_blocking(move || {
    let result = execute_process(&formated_config, |iteration, avg, best| {
            let progress = iteration as f64;
            let msg = serde_json::json!({
                "assigment_id": 0,
                "bestGeneration": [],
                "bestFitness": 0,
                "iteration": progress,
                "time": 0.0
            });
            println!("RESULT INT: {:?}", msg);
            let _ = tx.blocking_send(msg.to_string());
        });
    let duration = start.elapsed();
    let seconds = duration.as_secs() as i64;
    println!("RESULT: {:?}", result);
    let msg = serde_json::json!({
            "assigment_id": result.assigment_id,
            "bestGeneration": result.bestGeneration,
            "bestFitness": result.bestFitness,
            "iteration": result.iteration,
            "time": seconds
        });
    println!("RESULT FINAL: {:?}", msg);
    let _ = tx.blocking_send(msg.to_string());
    });
    // result.time = seconds;
    // result.assigment_id = assignment_id;
    // println!("RESULTADO ALGORITMO: \n{:#?}", result);
    // Json(result)
        EventStream! {
        while let Some(msg) = rx.recv().await {
            yield Event::data(msg);
        }
    }
}
