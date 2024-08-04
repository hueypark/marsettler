mod snake;

use avian2d::prelude::*;
use bevy::{audio::AudioPlugin, prelude::*, window::WindowPlugin};
use snake::{move_snakes, spawn_snake_head, SnakeHead};

fn main() {
    App::new()
        .add_plugins((
            DefaultPlugins
                .build()
                .set(WindowPlugin {
                    primary_window: Some(Window {
                        title: "Snake!".to_string(),
                        resolution: (1024.0, 768.0).into(),
                        ..default()
                    }),
                    ..default()
                })
                .disable::<AudioPlugin>(),
            PhysicsPlugins::default(),
        ))
        .insert_resource(PrintDebugTimer(Timer::from_seconds(
            1.0,
            TimerMode::Repeating,
        )))
        .insert_resource(Gravity(Vec2::ZERO))
        .add_systems(Startup, setup_camera)
        .add_systems(Startup, spawn_snake_head)
        .add_systems(Update, move_snakes)
        .add_systems(Update, print_debug_message)
        .run();
}

fn setup_camera(mut commands: Commands) {
    commands.spawn(Camera2dBundle::default());
}

#[derive(Resource)]
struct PrintDebugTimer(Timer);

fn print_debug_message(
    time: Res<Time>,
    mut timer: ResMut<PrintDebugTimer>,
    query: Query<&Transform, With<SnakeHead>>,
) {
    if !timer.0.tick(time.delta()).just_finished() {
        return;
    }

    for transform in query.iter() {
        info!("Snake head position: {:?}", transform.translation);
    }
}
