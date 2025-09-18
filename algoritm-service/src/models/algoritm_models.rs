use crate::models::{algoritm_config::AlgorithmConfig, group::Group, subject::Subject};

#[derive(Clone, Debug, PartialEq, PartialOrd)]
pub struct ClaseProgramada {
    pub group_id: usize,
    pub curso_id: usize,
    pub salon_id: usize,
    pub teacher_id: usize,
    pub dia: usize,
    pub bloque: usize,
}

pub struct HorarioBuilder<'a> {
    pub subject: Vec<Subject>,
    pub group: Vec<Group>,
    pub config: &'a AlgorithmConfig,
}

pub struct HorarioMutator {
    pub config: AlgorithmConfig,
}
