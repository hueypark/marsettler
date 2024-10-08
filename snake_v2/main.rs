mod leaderboard;
mod snake;
mod status;
mod ui;

use avian2d::prelude::*;
use bevy::{
    audio::AudioPlugin, log::LogPlugin, prelude::*, window::PrimaryWindow, window::WindowPlugin,
};
use bevy_jornet::JornetPlugin;
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
                .set(LogPlugin {
                    level: option_env!("BEVY_LOG_LEVEL")
                        .map(|s| s.parse().expect("Invalid log level"))
                        .unwrap_or(bevy::log::Level::INFO),
                    ..default()
                })
                .disable::<AudioPlugin>(),
            PhysicsPlugins::default(),
            JornetPlugin::with_leaderboard(
                option_env!("JORNET_LEADERBOARD_ID").unwrap_or(""),
                option_env!("JORNET_LEADERBOARD_KEY").unwrap_or(""),
            ),
        ))
        .insert_resource(Gravity(Vec2::ZERO))
        .insert_resource(status::Status::new())
        .add_event::<leaderboard::SendScoreEvent>()
        .add_systems(
            Startup,
            (
                setup_camera,
                setup_debug,
                spawn_snake_head,
                leaderboard::setup_leaderboard,
                ui::spawn_leaderboard,
            ),
        )
        .add_systems(
            Update,
            (
                rotate_snakes,
                move_snakes,
                print_cursor_world_position,
                on_esc_button,
                status::update_game_time,
                leaderboard::send_score,
                ui::update_status,
                ui::update_leaderboard,
            ),
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

fn on_esc_button(
    buttons: Res<ButtonInput<KeyCode>>,
    st: Res<status::Status>,
    mut event_writer: EventWriter<leaderboard::SendScoreEvent>,
) {
    if !buttons.just_pressed(KeyCode::Escape) {
        return;
    }

    event_writer.send(leaderboard::SendScoreEvent { score: st.score() });
}
