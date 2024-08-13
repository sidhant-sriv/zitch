let entities1 = [];
let entities2 = [];
let collisions = [];
let segmentDuration;
let totalTime;
let numSegments = 100; // Divide time into 100 segments
let currentTime = 0;

function setup() {
  createCanvas(1600, 800);
  totalTime = 10000; // Total time of simulation in milliseconds
  segmentDuration = totalTime / numSegments;

  // Create Type 1 Entities
  for (let i = 0; i < 100; i++) {
    entities1.push(new Entity(random(width), random(height), random(2, 5), true));
  }

  // Create Type 2 Entities.
  for (let i = 0; i < 100; i++) {
    entities2.push(new Entity(random(width), random(height), random(2, 5), false));
  }
}

function draw() {
  background(255);
  currentTime += deltaTime;

  // Update and display entities
  for (let entity of entities1) {
    entity.update();
    entity.display();
  }
  
  for (let entity of entities2) {
    entity.update();
    entity.display();
  }

  // Check for collisions
  for (let e1 of entities1) {
    for (let e2 of entities2) {
      if (e1.checkCollision(e2)) {
        collisions.push({
          time: currentTime,
          position: { x: e2.x, y: e2.y },
          segment: floor(currentTime / segmentDuration)
        });
      }
    }
  }

  // Display collisions for the current segment
  let segmentStart = floor(currentTime / segmentDuration) * segmentDuration;
  let segmentEnd = segmentStart + segmentDuration;

  for (let collision of collisions) {
    if (collision.time >= segmentStart && collision.time < segmentEnd) {
      fill(255, 0, 0);
      ellipse(collision.position.x, collision.position.y, 10, 10);
    }
  }

  // Stop the simulation after totalTime
  if (currentTime >= totalTime) {
    noLoop();
    saveCollisionData(); // Save collision data to JSON file
  }
}

// Entity class with a mix of random and deterministic behavior
class Entity {
  constructor(x, y, radius, isType1) {
    this.x = x;
    this.y = y;
    this.radius = radius;
    this.isType1 = isType1;
    
    // Random initial velocity and acceleration
    this.velocity = p5.Vector.random2D().mult(random(1, 3));
    this.acceleration = p5.Vector.random2D().mult(random(0.1, 0.5));

    // Add deterministic bias for type 1 entities
    if (this.isType1) {
      this.velocity.add(createVector(random(-0.5, 0.5), random(-0.5, 0.5)));
      this.acceleration.add(createVector(0.01, 0.01));
    }
  }

  update() {
    // Slight randomness in the update for all entities
    this.velocity.add(p5.Vector.random2D().mult(0.1));
    this.velocity.add(this.acceleration);
    this.x += this.velocity.x;
    this.y += this.velocity.y;

    // Bounce off edges with some damping
    if (this.x < 0 || this.x > width) {
      this.velocity.x *= -1;
      this.velocity.mult(0.99); // Damping effect
    }
    if (this.y < 0 || this.y > height) {
      this.velocity.y *= -1;
      this.velocity.mult(0.99); // Damping effect
    }
  }

  display() {
    fill(this.isType1 ? 0 : 100); // Different colors for type 1 and type 2
    ellipse(this.x, this.y, this.radius * 2);
  }

  checkCollision(other) {
    let d = dist(this.x, this.y, other.x, other.y);
    return d < this.radius + other.radius;
  }
}

// Function to save collision data to a JSON file
function saveCollisionData() {
  saveJSON(collisions, 'collisions.json');
}
    