use genevo::prelude::*;
use genevo::termination::{StopFlag, Termination, StopReason};
use genevo::simulation::State;


/// Nuestra struct de terminación
pub struct StagnationLimit {
    max_stagnant_generations: usize,
    last_best_fitness: usize,
    stagnant_generations: usize,
}

impl StagnationLimit {
    pub fn new(max_stagnant_generations: usize) -> Self {
        Self {
            max_stagnant_generations,
            last_best_fitness: 0,
            stagnant_generations: 0,
        }
    }
}

impl<A> Termination<A> for StagnationLimit
where
    A: Algorithm
{
    fn evaluate(&mut self, state: &State<A>) -> StopFlag {
        // let best_fitness = state.best_solution.solution.fitness;
        println!("RESULT:{:?}", state.result);
        let best_fitness = 15;

        if best_fitness > self.last_best_fitness {
            // hubo mejora → reseteamos
            self.last_best_fitness = best_fitness;
            self.stagnant_generations = 0;
            StopFlag::Continue
        } else {
            // no mejoró
            self.stagnant_generations += 1;
            if self.stagnant_generations >= self.max_stagnant_generations {
                StopFlag::StopNow("Simulation stopped after the maximum of 100 generations have been processed".to_string())
            } else {
                StopFlag::Continue
            }
        }
    }

    fn reset(&mut self) {
        self.last_best_fitness = 0;
        self.stagnant_generations = 0;
    }
}
