mod snake;

use avian2d::prelude::*;
use bevy::{audio::AudioPlugin, prelude::*, window::PrimaryWindow, window::WindowPlugin};
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
        .add_systems(
            Update,
            (
                move_snakes,
                print_debug_message,
                print_cursor_world_position,
            ),
        )
        .run();
}

#[derive(Component)]
struct MainCamera;

fn setup_camera(mut commands: Commands) {
    commands.spawn((Camera2dBundle::default(), MainCamera));
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

fn print_cursor_world_position(
    primary_query: Query<&Window, With<PrimaryWindow>>,
    camera_query: Query<(&Camera, &GlobalTransform), With<MainCamera>>,
    buttons: Res<ButtonInput<MouseButton>>,
) {
    if !buttons.just_pressed(MouseButton::Left) {
        return;
    }

    let Ok(window) = primary_query.get_single() else {
        return;
    };

    let Ok((camera, camera_transform)) = camera_query.get_single() else {
        return;
    };

    let Some(cursor_position) = window.cursor_position() else {
        return;
    };

    let Some(ray) = camera.viewport_to_world(camera_transform, cursor_position) else {
        return;
    };

    let world_position = ray.origin.truncate();

    info!(
        "Cursor world position: x: {}, y: {}",
        world_position.x, world_position.y
    );
}
