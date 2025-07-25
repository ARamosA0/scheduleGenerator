use serde::{Deserialize, Serialize};
use serde_json::Value;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AlgorithmConfig {
    pub num_cursos: usize,
    pub num_salones: usize,
    pub num_dias: usize,
    pub num_bloques: usize,
    // pub population_size: usize,
    // pub generation_limit: u64,
    // pub selection_ratio: f64,
    // pub mutation_rate: f64,
    // pub reinsertion_ratio: f64,
    pub population: usize,
    pub generations: u64,
    pub mutation: f64,
    pub cross_over: f64,
    pub reinsertion: f64,
    pub elitism: f64,

    pub template: Value,
    pub subjects: Value,
    pub teachers: Value,
    pub rooms: Value,
}
