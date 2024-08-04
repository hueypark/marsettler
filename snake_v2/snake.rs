use avian2d::prelude::*;
use bevy::prelude::*;

pub fn spawn_snake_head(mut cmds: Commands) {
    const LENGTH: f32 = 10.0;

    cmds.spawn((
        SpriteBundle {
            sprite: Sprite {
                color: SNAKE_HEAD_COLOR,
                ..default()
            },
            transform: Transform {
                scale: Vec3::splat(LENGTH),
                ..default()
            },
            ..default()
        },
        RigidBody::Dynamic,
        Collider::rectangle(LENGTH, LENGTH),
        SnakeHead {},
    ));
}

pub fn move_snakes(mut query: Query<(&mut LinearVelocity, &Transform), With<SnakeHead>>) {
    for (mut vel, transform) in query.iter_mut() {
        let forward = transform.rotation * Vec3::Y;
        vel.x = forward.x * 100.0;
        vel.y = forward.y * 100.0;
    }
}

#[derive(Component)]
pub struct SnakeHead {}

const SNAKE_HEAD_COLOR: Color = Color::srgb(0.7, 0.7, 0.7);
