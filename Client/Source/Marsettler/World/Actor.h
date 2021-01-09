#pragma once

class Actor
{
public:
	// 생성자
	Actor(const int64 id, const FVector location);

	// 위치 갱신 여부를 반환한다
	bool GetLocationUpdated() const;

	// ID를 반환한다.
	int64 ID() const;

	// 위치를 반환한다.
	FVector Location() const;

	// 위치를 설정한다.
	void SetLocation(const FVector location);

	// 위치 갱신 여부를 반환한다
	void SetLocationUpdated(const bool updated);

private:
	// ID
	int64 m_id;

	// 위치
	FVector m_location;

	// 위치 갱신 여부
	bool m_locationUpdated;
};
