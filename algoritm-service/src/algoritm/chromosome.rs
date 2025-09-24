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
        let mut genes = Vec::new();
        for subject in &self.subject {
            for _ in 0..subject.hours {
                let group = self.group.get(0).expect("No se encontro grupo");
                let subject_id: usize =
                    create_subject_by_group(group).expect("Grupo sin cursos");
                let room_id: usize =
                    create_room_by_subject(&self.config.subjects, subject_id).expect("Room invalida");
                let teacher_id: usize =
                    create_teacher_by_subject(&self.config.teachers, &self.config.subjects, subject_id).expect("teacher invalido");

                // println!("SUBJECT: {:?}", subject);

                genes.push(ClaseProgramada {
                    group_id: group.id,
                    curso_id: subject_id,
                    salon_id: room_id,
                    teacher_id: teacher_id,
                    dia: random_day(rng, self.config.num_days),
                    bloque: random_block(rng, self.config.num_periods),
                });
            }
        }
        // println!("GENES:{:?}", genes);
        genes
    }
}

fn random_teacher<R: Rng + ?Sized>(rng: &mut R, max: usize) -> usize {
    rng.gen_range(1..max)
}

fn random_day<R: Rng + ?Sized>(rng: &mut R, max: usize) -> usize {
    rng.gen_range(0..max)
}

fn random_block<R: Rng + ?Sized>(rng: &mut R, max: usize) -> usize {
    rng.gen_range(0..max)
}

fn create_subject_by_group(group: &Group) -> Option<usize> {
    let mut rng = thread_rng();
    group.subjects.choose(&mut rng).copied()
}



fn create_room_by_subject(subjects: &Vec<Subject>, subject_id: usize) -> Option<usize> {
    if let Some(subject) = subjects.iter().find(|s| s.id == subject_id) {
        Some(subject.required_room_type)
    } else {
        None
    }

}

fn create_teacher_by_subject(teachers: &Vec<Teacher>, subjects: &Vec<Subject>, subject_id: usize) -> Option<usize> {
    if let Some(subject) = subjects.iter().find(|s| s.id == subject_id) {
        let mut candidates: Vec<&Teacher> = teachers
            .iter()
            .filter(|t| t.speciality == subject.speciality)
            .collect();

        if candidates.is_empty() {
            return None;
        }

        let mut rng = rand::thread_rng();
        candidates
            .choose(&mut rng)
            .map(|teacher| teacher.id)
    } else {
        None
    }
}

