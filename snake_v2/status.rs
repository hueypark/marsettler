use bevy::prelude::*;
use bevy::time::Stopwatch;

#[derive(Resource)]
pub struct Status {
    stopwatch: Stopwatch,
}

impl Status {
    pub fn new() -> Self {
        Self {
            stopwatch: Stopwatch::new(),
        }
    }

    pub fn time(&self) -> f32 {
        return self.stopwatch.elapsed().as_secs_f32();
    }
}

pub fn update_game_time(mut st: ResMut<Status>, time: Res<Time>) {
    st.stopwatch.tick(time.delta());
}
