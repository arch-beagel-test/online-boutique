"""Loyalty points shown in the order confirmation email."""

POINTS_PER_UNIT = 10


def points_for(units):
    return max(0, int(units)) * POINTS_PER_UNIT
