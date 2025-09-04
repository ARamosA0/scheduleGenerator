use crate::algoritm::schedule_main::{HorarioGenome, HorarioGenomeDisplay, HorarioWrapper};
use crate::models::{algoritm_config::AlgorithmConfig, algoritm_models::*, subject::Subject, group::Group};
use genevo::genetic::{Parents, ParentsSlice};
use rocket::{config, serde::json::Json};

use crate::algoritm::demo::{run_genetic_algorithm, BestSolutionResult};
use genevo::operator::CrossoverOp;
use genevo::recombination::discrete::MultiPointCrossBreeder;
use genevo::{operator::prelude::*, prelude::*, random::Rng, types::fmt::Display};
use genevo::prelude::Genotype;
use serde::Serialize;
use rand::seq::SliceRandom;
use rand::thread_rng;


// impl MultiPointCrossover for HorarioWrapper {
//     type Dna = i32;
//     fn crossover<R>(parents: Parents<Self>, num_cut_points: usize, rng: &mut R)  -> Children<Self>
//     where
//         R: Rng + Sized,
//     {
        
//         let (child1, child2) = multi_point_crossover(parents, num_cut_points, rng);
//         (HorarioWrapper(child1), HorarioWrapper(child2))

//     }
// }

// fn multi_point_crossover<R: Rng>(
//     parent1: &HorarioGenome,
//     parent2: &HorarioGenome,
//     num_points: usize,
//     rng: &mut R,
// ) -> (HorarioGenome, HorarioGenome) {
//     let len = parent1.len();
//     assert_eq!(len, parent2.len());
    
//     // Generar puntos de corte aleatorios
//     let mut points: Vec<usize> = (0..num_points)
//         .map(|_| rng.gen_range(0..len))
//         .collect();
//     points.sort();
//     points.dedup();
    
//     let mut child1 = Vec::with_capacity(len);
//     let mut child2 = Vec::with_capacity(len);
    
//     let mut current = 0;
//     let mut use_parent1 = true;
    
//     for &point in &points {
//         while current < point {
//             if use_parent1 {
//                 child1.push(parent1[current].clone());
//                 child2.push(parent2[current].clone());
//             } else {
//                 child1.push(parent2[current].clone());
//                 child2.push(parent1[current].clone());
//             }
//             current += 1;
//         }
//         use_parent1 = !use_parent1;
//     }
    
//     // Añadir el resto
//     while current < len {
//         if use_parent1 {
//             child1.push(parent1[current].clone());
//             child2.push(parent2[current].clone());
//         } else {
//             child1.push(parent2[current].clone());
//             child2.push(parent1[current].clone());
//         }
//         current += 1;
//     }
    
//     (child1, child2)
// }

// pub struct CustomMultiPointCrossover {
//     num_points: usize,
// }

// impl CustomMultiPointCrossover {
//     pub fn new(num_points: usize) -> Self {
//         Self { num_points }
//     }
// }

// impl CrossBreeder<HorarioWrapper> for CustomMultiPointCrossover {
//     fn cross_breed<R>(
//         &self,
//         parents: Parents<HorarioWrapper>,
//         rng: &mut R,
//     ) -> Children<HorarioWrapper>
//     where
//         R: Rng + Sized,
//     {
//         let mut children = Vec::new();
//         for pair in parents.chunks(2) {
//             if pair.len() == 2 {
//                 let (child1, child2) = pair[0].crossover(&pair[1], self.num_points, rng);
//                 children.push(child1);
//                 children.push(child2);
//             }
//         }
//         children
//     }
// }
// pub struct CustomCrossover;
// impl CrossoverOp<G> for CustomCrossover{
//     fn crossover<R>(value: Self, parents: &Self, rng: &R) -> Self
//     where
//         R: Rng + Sized,
//     {
//         // ClaseProgramada
//     }
// }

