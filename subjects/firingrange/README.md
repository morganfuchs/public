## FiringRange

In this exercise, you will learn to use Unreal Engine and use Blueprints to script functionalities of a basic FPS game. The theme is to reproduce a firing range. And remember weapons are allowed only inside the shooting area!!

### Instructions

The map of this project should be composed by a cube with dimensions of X = 35, Y = 40, Z = 1, simulating the floor and other cubes as walls around the floor. In the map there should be a zone where the character should be able to walk around and shoot to the targets and another zone where the player can not go, where the targets are present.

You will need this [image](), from which you will create a new material and associate the image to it.

For this project you will have to create a Blueprint Class target, that will have some characteristics. The target should :

- have associated to it the previous created material.
- either be moving from side to side or stationary.
- be dynamic, using the timeline node.
- use a public variable to set or unset the movement animation of the target.
- rise again after x seconds after being hit it with a projectile, , and behave like before it was hit.

Only one class of target is allowed on the whole project.

The previous mentioned projectile should:

- have a size of X = Y = Z = 0,5.
- have a speed of 10000.

When finished, your project should look like the executable example on the folder or the ["Expected Result" video](https://youtu.be/EBibaN-dh_0).

> Don’t forget to zip up the project compile and save everything for peer correction.

#### Bonus

Here are some ideas for improving the game:

- Targets with different speeds
- Textures on the walls and ground
- Add obstacles in front of the targets