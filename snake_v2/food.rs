use avian2d::prelude::*;
use bevy::prelude::*;

pub fn spawn_food(mut cmds: Commands, position: Vec2) {
    const LENGTH: f32 = 10.0;
    const FOOD_COLOR: Color = Color::srgb(0.0, 0.7, 0.0);

    cmds.spawn((
        Sprite {
            color: FOOD_COLOR,
            custom_size: Some(Vec2::splat(LENGTH)),
            ..default()
        },
        Transform {
            scale: Vec3::splat(LENGTH),
            translation: Vec3::new(position.x, position.y, 0.0),
            ..default()
        },
        RigidBody::Static,
        Collider::rectangle(LENGTH, LENGTH),
    ));
}
