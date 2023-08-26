from typing import Any


def generate_all_subsequences(sequence: list[Any]) -> None:
    create_state_space_tree(sequence, [], 0)


def create_state_space_tree(
    sequence: list[Any], current_subsequence: list[Any], index: int
) -> None:
    """
    Creates a state space tree to iterate through each branch using DFS.
    We know that each state has exactly two children.
    It terminates when it reaches the end of the given sequence.
    """

    if index == len(sequence):
        print(current_subsequence)
        return

    create_state_space_tree(sequence, current_subsequence, index + 1)
    current_subsequence.append(sequence[index])
    create_state_space_tree(sequence, current_subsequence, index + 1)
    current_subsequence.pop()


if __name__ == "__main__":
    seq: list[Any] = [3, 1, 2, 4]
    #generate_all_subsequences(seq)

    seq.clear()
    seq.extend(["A", "B", "C", "D"])
    generate_all_subsequences(seq)
