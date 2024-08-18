use crate::status;
use bevy::prelude::*;
use bevy_jornet::Leaderboard;
use std::cmp::Ordering;

#[derive(Component)]
pub struct StatusUI;

#[derive(Component)]
pub struct LeaderboardUI;

const BACKGROUND_COLOR: Color = Color::srgba(0.1, 0.1, 0.1, 0.80);

pub fn spawn_leaderboard(mut commands: Commands) {
    commands
        .spawn(NodeBundle {
            style: Style {
                margin: UiRect::new(Val::Auto, Val::Px(15.0), Val::Px(10.0), Val::Auto),
                justify_content: JustifyContent::Start,
                align_items: AlignItems::Start,
                flex_direction: FlexDirection::Row,
                ..default()
            },
            ..default()
        })
        .with_children(|parent| {
            parent
                .spawn(NodeBundle {
                    style: Style {
                        margin: UiRect::new(Val::Px(15.0), Val::Auto, Val::Px(10.0), Val::Auto),
                        justify_content: JustifyContent::Start,
                        align_items: AlignItems::Start,
                        flex_direction: FlexDirection::Column,
                        ..default()
                    },
                    background_color: BACKGROUND_COLOR.into(),
                    ..default()
                })
                .with_children(|parent02| {
                    parent02.spawn(
                        TextBundle::from_section(
                            "Leaderboard",
                            TextStyle {
                                font_size: 22.0,
                                ..default()
                            },
                        )
                        .with_style(Style {
                            margin: UiRect::px(10.0, 10.0, 5.0, 5.0),
                            ..default()
                        }),
                    );

                    parent02
                        .spawn(NodeBundle {
                            style: Style {
                                flex_direction: FlexDirection::Column,
                                justify_content: JustifyContent::Start,
                                align_items: AlignItems::Start,
                                margin: UiRect::px(5.0, 10.0, 5.0, 5.0),
                                ..default()
                            },
                            ..default()
                        })
                        .insert(LeaderboardUI);
                });

            parent
                .spawn(NodeBundle {
                    style: Style {
                        margin: UiRect::new(Val::Px(15.0), Val::Auto, Val::Px(10.0), Val::Auto),
                        justify_content: JustifyContent::Start,
                        align_items: AlignItems::Start,
                        flex_direction: FlexDirection::Column,
                        ..default()
                    },
                    background_color: BACKGROUND_COLOR.into(),
                    ..default()
                })
                .with_children(|parent02| {
                    parent02.spawn(
                        TextBundle::from_section(
                            "Status",
                            TextStyle {
                                font_size: 22.0,
                                ..default()
                            },
                        )
                        .with_style(Style {
                            margin: UiRect::px(10.0, 10.0, 5.0, 5.0),
                            ..default()
                        }),
                    );

                    parent02.spawn((
                        TextBundle::from_section(
                            "Time: 0s",
                            TextStyle {
                                font_size: 20.0,
                                ..default()
                            },
                        )
                        .with_style(Style {
                            margin: UiRect::px(10.0, 10.0, 5.0, 5.0),
                            ..default()
                        }),
                        StatusUI,
                    ));
                });
        });
}

pub fn update_status(st: Res<status::Status>, mut query: Query<&mut Text, With<StatusUI>>) {
    for mut text in query.iter_mut() {
        text.sections[0].value = format!("Time: {:.2}s", st.time());
    }
}

pub fn update_leaderboard(
    leaderboard: Res<Leaderboard>,
    mut commands: Commands,
    query: Query<Entity, With<LeaderboardUI>>,
) {
    if !leaderboard.is_changed() {
        return;
    }

    let mut scores = leaderboard.get_leaderboard();
    scores.sort_by(|s1, s2| s2.score.partial_cmp(&s1.score).unwrap_or(Ordering::Equal));

    let ui = query.single();
    commands.entity(ui).despawn_descendants();

    for score in scores {
        commands.entity(ui).with_children(|parent| {
            parent.spawn(TextBundle::from_section(
                format!("{}: {:.2}", score.player, score.score),
                TextStyle {
                    font_size: 20.0,
                    color: Color::WHITE,
                    ..default()
                },
            ));
        });
    }
}
