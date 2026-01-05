from dataclasses import dataclass
from scipy.optimize import milp, LinearConstraint, Bounds
import numpy as np


@dataclass
class Machine:
    lights: int
    buttons: list[int]
    jolts: list[int]


def read_inputs(path: str) -> list[Machine]:
    with open(path) as f:
        lines = f.read().strip().split("\n")

    machines = []
    for line in lines:
        s_lights = 1
        e_lights = line.index("]") - 1
        s_jolts = line.index("{")
        e_jolts = line.index("}")
        e_buttons = s_jolts - 2
        s_buttons = e_lights + 2

        lights = 0
        for i in range(s_lights, e_lights + 1):
            if line[i] == "#":
                lights |= 1 << (i - s_lights)

        jolts_line = line[s_jolts + 1 : e_jolts]
        jolts = [int(v) for v in jolts_line.split(",")]

        buttons = []
        button_line = line[s_buttons : e_buttons + 1]
        button_vals = button_line.split()
        for v in button_vals:
            inner = v[1:-1]
            mask = 0
            for num in inner.split(","):
                mask |= 1 << int(num)
            buttons.append(mask)

        machines.append(Machine(lights=lights, buttons=buttons, jolts=jolts))

    return machines


def solve_machine(buttons: list[int], jolts: list[int]) -> int:
    n_counters = len(jolts)
    n_buttons = len(buttons)

    A = np.zeros((n_counters, n_buttons), dtype=float)
    for btn_idx, mask in enumerate(buttons):
        for counter in range(n_counters):
            if mask & (1 << counter):
                A[counter][btn_idx] = 1

    b = np.array(jolts, dtype=float)

    c = np.ones(n_buttons)

    constraints = LinearConstraint(A, b, b)

    bounds = Bounds(lb=0, ub=np.inf)
    integrality = np.ones(n_buttons)

    result = milp(c, constraints=constraints, bounds=bounds, integrality=integrality)

    return int(round(result.fun))


def part2():
    machines = read_inputs("inputs/input-10.txt")
    total = 0
    for m in machines:
        print(f"Machine: {m}")
        presses = solve_machine(m.buttons, m.jolts)
        total += presses

    print(f"Total: {total}")


if __name__ == "__main__":
    part2()
