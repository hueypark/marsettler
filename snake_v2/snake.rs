use bevy::prelude::*;

pub fn spawn_snake_head(mut cmds: Commands) {
    cmds.spawn(SpriteBundle {
        sprite: Sprite {
            color: SNAKE_HEAD_COLOR,
            ..default()
        },
        transform: Transform {
            scale: Vec3::splat(10.0),
            ..default()
        },
        ..default()
    })
    .insert(SnakeHead {});
}

pub fn move_snakes(mut query: Query<&mut Transform, With<SnakeHead>>) {
    for mut transform in query.iter_mut() {
        transform.translation.x += 1.0;
    }
}

#[derive(Component)]
pub struct SnakeHead {}

const SNAKE_HEAD_COLOR: Color = Color::srgb(0.7, 0.7, 0.7);
