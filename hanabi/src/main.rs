mod camera;
mod effect;

use bevy::prelude::*;
use bevy_hanabi::prelude::*;

fn main() {
    App::new()
        .add_plugins((DefaultPlugins, HanabiPlugin))
        .add_systems(
            Startup,
            (camera::setup, effect::setup_blacksmith_hammer_sparks),
        )
        .add_systems(Update, (effect::update_spawn, effect::update))
        .run();
}
