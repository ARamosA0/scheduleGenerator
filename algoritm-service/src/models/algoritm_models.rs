use crate::models::{algoritm_config::AlgorithmConfig, subject::Subject};

#[derive(Clone, Debug, PartialEq, PartialOrd)]
pub struct ClaseProgramada {
    pub curso_id: usize,
    pub salon_id: usize,
    pub dia: usize,
    pub bloque: usize,
}

pub struct HorarioBuilder {
    pub subject: Vec<Subject>,
    pub config: AlgorithmConfig,
}

pub struct HorarioMutator {
    pub config: AlgorithmConfig,
}
