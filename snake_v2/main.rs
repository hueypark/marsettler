mod snake;

use avian2d::prelude::*;
use bevy::{audio::AudioPlugin, prelude::*, window::PrimaryWindow, window::WindowPlugin};
use snake::{move_snakes, rotate_snakes, spawn_snake_head, SnakeHead};

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
        .insert_resource(Gravity(Vec2::ZERO))
        .add_systems(Startup, (setup_camera, setup_debug, spawn_snake_head))
        .add_systems(
            Update,
            (rotate_snakes, move_snakes, print_cursor_world_position),
        )
        .run();
}

#[derive(Component)]
struct MainCamera;

fn setup_camera(mut commands: Commands) {
    commands.spawn((Camera2dBundle::default(), MainCamera));
}

fn setup_debug(mut commands: Commands) {
    commands.spawn((TextBundle::from_section(
        String::from("v") + env!("CARGO_PKG_VERSION"),
        TextStyle {
            font_size: 40.0,
            ..default()
        },
    )
    .with_text_justify(JustifyText::Center)
    .with_style(Style {
        position_type: PositionType::Absolute,
        top: Val::Px(5.0),
        left: Val::Px(5.0),
        ..default()
    }),));
}

fn print_cursor_world_position(
    primary_query: Query<&Window, With<PrimaryWindow>>,
    camera_query: Query<(&Camera, &GlobalTransform), With<MainCamera>>,
    mut snake_head_query: Query<&mut SnakeHead, With<SnakeHead>>,
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

    let world_position = Vec3::new(ray.origin.x, ray.origin.y, 0.0);

    for mut snake_head in snake_head_query.iter_mut() {
        snake_head.desired_position = world_position;
    }
}
