use crate::models::{algoritm_config::AlgorithmConfig, algoritm_models::*, group::Group, subject::Subject, room::Room, teacher::Teacher};
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use rand::seq::SliceRandom;
use rand::thread_rng;

type HorarioGenome = Vec<ClaseProgramada>;

impl<'a> GenomeBuilder<HorarioGenome> for HorarioBuilder<'a> {
    fn build_genome<R>(&self, _: usize, rng: &mut R) -> HorarioGenome
    where
        R: Rng + Sized,
    {

        self.subject
            .iter()
            .map(|subject: &Subject| {

            let group = self.group.get(0).expect("No se encontro grupo");
            let subject_id: usize= create_subject_by_group(group).expect("Grupo sin cursos");
            let room_id:usize = create_room_by_subject(&self.config.subjects, subject_id).expect("Room invalida");
            let teacher_id: usize = create_teacher_by_subject(&self.config.subjects).expect("teacher invalido");
            ClaseProgramada {
                group_id: group.id,
                curso_id: subject_id,
                salon_id: room_id,
                teacher_id: rng.gen_range(0..self.config.num_teachers),
                dia: rng.gen_range(0..self.config.num_days),
                bloque: rng.gen_range(0..self.config.num_periods),
            }
            })
            .collect()

    }
}

fn create_subject_by_group(group: &Group) -> Option<usize> {
    let mut rng = thread_rng();
    group.subjects.choose(&mut rng).copied()
}



fn create_room_by_subject(subjects: &Vec<Subject>, subject_id: usize) -> Option<usize> {
    let mut rng= thread_rng();
    if let Some(subject) = subjects.iter().find(|s| s.id == subject_id) {
        Some(subject.required_room_type)
    } else {
        None
    }

}

fn create_teacher_by_subject(subject: &Vec<Subject>) -> Option<usize> {
    
    Some(0)
}

