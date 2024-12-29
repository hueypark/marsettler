use crate::status;
use bevy::prelude::*;

#[derive(Component)]
pub struct StatusUI;

pub fn update_status(
    st: Res<status::Status>,
    mut query: Query<Entity, With<StatusUI>>,
    mut writer: TextUiWriter,
) {
    for text in query.iter_mut() {
        *writer.text(text, 0) = format!("Time: {:.2}s", st.time());
    }
}
