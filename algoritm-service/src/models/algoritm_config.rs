use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AlgorithmConfig {
    pub num_profesores: usize,
    pub num_cursos: usize,
    pub num_salones: usize,
    pub num_dias: usize,
    pub num_bloques: usize,
    pub population_size: usize,
    pub generation_limit: u64,
    pub selection_ratio: f64,
    pub mutation_rate: f64,
    pub reinsertion_ratio: f64,
}
